package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var btnConnect *widget.Button
var connectContainer *fyne.Container
var actionConnect = true

func getConnectComponents() *fyne.Container {
	if btnConnect == nil {
		btnConnect = widget.NewButton(
			locales.Text("gen9"),
			func() {

			},
		)
		connectContainer = container.NewCenter(btnConnect)
	}
	return connectContainer
}

func updateConnectButtonStatus() {
	if btnConnect == nil {
		return
	}
	switch cg.GetCurrentState() {

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

	btnConnect.Refresh()
}

func disableForm() {
	// btnDelProfile.Disable()
	// btnSaveProfile.Disable()
	// selectProfile.Disable()
	// selectCity.Disable()
	// selectCountry.Disable()
	// selectServerType.Disable()
	// selectServerInstance.Disable()
	// selectConnection.Disable()
	// selectService.Disable()
	_enableForm(false)
}

func enableForm() {
	// if len(selectProfile.Selected) > 0 {
	// 	btnDelProfile.Enable()
	// }
	// btnSaveProfile.Enable()
	// selectProfile.Enable()
	// selectCity.Enable()
	// selectCountry.Enable()
	// selectServerType.Enable()
	// selectServerInstance.Enable()
	// selectConnection.Enable()
	// selectService.Enable()
	_enableForm(true)
}

func _enableForm(enable bool) {
	components := []fyne.Disableable{
		btnDelProfile,
		btnSaveProfile,
		selectProfile,
		selectCity,
		selectCountry,
		selectServerType,
		selectServerInstance,
		selectConnection,
		selectService,
	}
	for _, c := range components {
		if c != nil {
			switch enable {
			case true:
				if c.Disabled() {
					if c == selectProfile {
						if len(selectProfile.Selected) > 0 {
							c.Enable()
						}
					} else {
						fmt.Printf("Enabling %p\n", c)
						c.Enable()
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
