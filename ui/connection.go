package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"

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
	btnDelProfile.Disable()
	btnSaveProfile.Disable()
	selectProfile.Disable()
	selectCity.Disable()
	selectCountry.Disable()
	selectServerType.Disable()
	selectServerInstance.Disable()
	selectConnection.Disable()
	selectService.Disable()
}

func enableForm() {
	if len(selectProfile.Selected) > 0 {
		btnDelProfile.Enable()
	}
	btnSaveProfile.Enable()
	selectProfile.Enable()
	selectCity.Enable()
	selectCountry.Enable()
	selectServerType.Enable()
	selectServerInstance.Enable()
	selectConnection.Enable()
	selectService.Enable()
}
