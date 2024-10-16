package ui

import (
	"cyberghostvpn-gui/about"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var mainApp fyne.App
var mainWindow fyne.Window

/* Public Functions */

func GetMainWindow() fyne.Window {

	if mainWindow == nil {
		mainWindow = getApp().NewWindow(fmt.Sprintf("%s - v%s", about.AppName, about.AppVersion))
		mainWindow.SetFixedSize(true)
		mainWindow.SetContent(getMainContent())
	}

	return mainWindow
}

/* Private Functions */

func getApp() fyne.App {
	if mainApp == nil {
		mainApp = app.New()
	}
	return mainApp
}

func getMainContent() *fyne.Container {
	text2 := canvas.NewText("2", color.White)
	text3 := canvas.NewText("3", color.White)
	vBox := layout.NewVBoxLayout()
	return container.New(vBox, getStatusBox(), text2, text3)
}
