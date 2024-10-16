package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var statusBox *fyne.Container
var textStatus *canvas.Text

type CustomColor color.Color

var StatusColorRed CustomColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var StatusColorGreen CustomColor = color.RGBA{R: 0, G: 255, B: 0, A: 255}
var StatusColorBlue CustomColor = color.RGBA{R: 0, G: 0, B: 255, A: 255}
var StatusColorOrange CustomColor = color.RGBA{R: 255, G: 128, B: 0, A: 255}

func getStatusBox() *fyne.Container {

	if statusBox == nil {
		// Layout
		mainBox := layout.NewHBoxLayout()

		// Status
		textDefStatus := canvas.NewText(locales.Text("sta0"), color.White)
		textStatus = canvas.NewText(locales.Text("sta6"), color.White)
		hBox := layout.NewHBoxLayout()
		statusContainer := container.New(hBox, textDefStatus, textStatus)

		// Create Status Box
		statusBox = container.New(mainBox, getCyberGhostLogo(), statusContainer)
	}

	return statusBox
}

func updateStatus() {
	status := cg.GetCurrentState()

	switch status {
	case cg.Connected:
		textStatus.Text = locales.Text("sta1")
		textStatus.Color = StatusColorGreen
	case cg.Disconnected:
		textStatus.Text = locales.Text("sta2")
		textStatus.Color = StatusColorRed
	case cg.Connecting:
		textStatus.Text = locales.Text("sta3")
		textStatus.Color = StatusColorOrange
	case cg.Disconnecting:
		textStatus.Text = locales.Text("sta4")
		textStatus.Color = StatusColorOrange
	case cg.NotInstalled:
		textStatus.Text = locales.Text("sta5")
		textStatus.Color = StatusColorRed
	default:
		textStatus.Text = locales.Text("sta6")
		textStatus.Color = color.White
	}

	textStatus.Refresh()
}
