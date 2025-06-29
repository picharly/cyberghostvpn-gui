package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/settings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var btnConnect *widget.Button
var connectContainer *fyne.Container
var actionConnect = true
var needPassword = false

// getConnectComponents initializes and returns a container with a connect button.
// The button toggles between connecting and disconnecting the VPN based on
// the current state. It uses ShowPopupSudo to execute the connect or disconnect
// command with administrative privileges.
func getConnectComponents() *fyne.Container {
	if btnConnect == nil {
		btnConnect = widget.NewButton(
			locales.Text("gen9"),
			func() {
				go hideAfterStatusChange()
				if actionConnect {
					ShowPopupSudo(cg.Connect()...)
				} else {
					ShowPopupSudo(cg.Disconnect()...)
				}
			},
		)
		connectContainer = container.NewCenter(btnConnect)
	}
	return connectContainer
}

// hideAfterStatusChange hides the main window after the VPN status changes or after 30 seconds maximum.
// It is used when the user clicks the "Connect" button to hide the window after the VPN is connected or disconnected.
func hideAfterStatusChange() {

	// Get current settings
	cfg, _ := settings.GetCurrentSettings()
	// Check if HideWhenConnected is enabled
	if !cfg.HideWhenConnected {
		return
	}

	currentState := cg.CurrentState
	start := time.Now()
	for {
		if cg.CurrentState != currentState {
			GetMainWindow().Hide()
			break
		}
		if time.Since(start).Seconds() > 30 {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
}

// updateConnectButtonStatus updates the connect button status based on the current
// state of the CyberGhost VPN. It is called when the status of the VPN changes.
// It disables the button if the country is not selected and the VPN is not connected.
// It displays the appropriate text on the button based on the current state
// of the VPN.
func updateConnectButtonStatus() {
	if btnConnect == nil {
		return
	}

	if len(cg.SelectedCountry.Name) == 0 && cg.CurrentState != cg.Connected {
		if !btnConnect.Disabled() {
			btnConnect.Disable()
		}
		return
	}

	switch cg.CurrentState {

	case cg.Connected:
		btnConnect.Text = locales.Text("gen14")
		actionConnect = false
		btnConnect.Enable()
		disableForm()
	case cg.Disconnected:
		btnConnect.Text = locales.Text("gen9")
		actionConnect = true
		btnConnect.Enable()
		enableForm()
	case cg.Unknown:
		btnConnect.Text = locales.Text("gen9")
		actionConnect = true
		btnConnect.Disable()
		disableForm()
	default:
		btnConnect.Text = locales.Text("gen9")
		actionConnect = true
		btnConnect.Disable()
		disableForm()
	}

	// go func() {
	fyne.DoAndWait(func() {
		btnConnect.Refresh()
	})
	// }()
}

// disableForm disables all the form elements so that the user can't make changes when the VPN is connected.
func disableForm() {
	_enableForm(false)
}

// enableForm enables all the form elements so that the user can make changes when the VPN is disconnected.
func enableForm() {
	_enableForm(true)
}

// _enableForm enables or disables all the form components depending on the value of the enable parameter.
// The components are:
// - the delete profile button
// - the save profile button
// - the profile select widget
// - the city select widget
// - the country select widget
// - the server type select widget
// - the server instance select widget
// - the connection select widget
// - the VPN service select widget
//
// The method iterates over the components and enables or disables them one by one.
// If the component is already enabled and the enable parameter is true, or if the component is disabled and the enable parameter is false, the method does nothing.
// If the enable parameter is true, the profile select widget is enabled only if a profile is currently selected.
// The method also checks if the select widget is not empty before enabling it.
func _enableForm(enable bool) {
	components := []fyne.Disableable{
		selectProfile,
		btnDelProfile,
		btnSaveProfile,
		selectCity,
		selectCountry,
		selectServerType,
		selectServerInstance,
		selectConnection,
		selectService,
		selectStreamingService,
	}
	for _, c := range components {
		if c != nil {
			switch enable {
			case true:
				if c.Disabled() {
					if c == btnDelProfile {
						if len(selectProfile.Selected) > 0 {
							c.Enable()
						}
					} else {
						if s, ok := c.(*widget.Select); ok {
							if len(s.Options) > 1 || s == selectProfile {
								s.Enable()
							}
						} else if b, ok := c.(*widget.Button); ok {
							b.Enable()
						}
					}
				}
			case false:
				if !c.Disabled() {
					c.Disable()
				}
			}
		}
	}
}

// _automaticEnableDisable enables or disables the given select widget automatically based on the number of options it has.
// If the select widget has less than 2 options, it is disabled. Otherwise, it is enabled.
// The method runs in a separate goroutine and checks the number of options every 100 milliseconds.
// The method is meant to be called once and is not meant to be called again until the select widget is disposed of.
func _automaticEnableDisable(selectComponent *widget.Select) {
	// Automatic Enable/Disable
	go func(s *widget.Select) {
		for {
			if cg.CurrentState != cg.Connected {
				if len(s.Options) < 2 {
					if !s.Disabled() {
						s.Disable()
					}
				} else {
					if s.Disabled() {
						s.Enable()
					}
				}
			}
			time.Sleep(time.Millisecond * 100)
		}
	}(selectComponent)
}
