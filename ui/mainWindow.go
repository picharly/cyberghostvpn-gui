package ui

import (
	"cyberghostvpn-gui/about"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var mainApp fyne.App
var mainWindow fyne.Window

/* Public Functions */

func GetMainWindow() fyne.Window {

	if mainWindow == nil {
		mainWindow = getApp().NewWindow(fmt.Sprintf("%s - v%s", about.AppName, about.AppVersion))
		mainWindow.SetContent(widget.NewLabel("Hello World!"))
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
