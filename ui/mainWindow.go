package ui

import (
	"cyberghostvpn-gui/about"
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/settings"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

var mainApp fyne.App
var mainWindow fyne.Window

/* Public Functions */

func GetApp() fyne.App {
	if mainApp == nil {
		mainApp = app.NewWithID("com.github.picharly.cyberghostvpn-gui")
		mainApp.Settings().SetTheme(&resources.DarkTheme{Theme: theme.DefaultTheme()})
		mainApp.SetIcon(resources.GetCyberGhostIcon())
	}
	return mainApp
}

func GetMainWindow() fyne.Window {

	if mainWindow == nil {
		mainWindow = GetApp().NewWindow(fmt.Sprintf("%s - v%s", about.AppName, about.AppVersion))
		mainWindow.SetFixedSize(true)
		mainWindow.SetContent(getMainContent())

		// Set tray icon
		cfg, _ := settings.GetCurrentSettings()
		if cfg.TrayIcon {
			setTrayIcon(mainWindow)
		}

	}

	return mainWindow
}

/* Private Functions */

func getMainContent() *fyne.Container {
	text2 := canvas.NewText("2", resources.ColorWhite)
	text3 := canvas.NewText("3", resources.ColorWhite)
	vBox := layout.NewVBoxLayout()
	return container.New(vBox, getInfoBox(), getTabs(), text2, text3)
}

func setTrayIcon(window fyne.Window) {
	if window == nil {
		return
	}
	if desk, ok := GetApp().(desktop.App); ok {
		m := fyne.NewMenu(about.AppName,
			fyne.NewMenuItem("Hide", func() { window.Hide() }),
			fyne.NewMenuItem("Show", func() { window.Show() }),
		)
		desk.SetSystemTrayMenu(m)
	}

	window.SetCloseIntercept(func() {
		cfg, _ := settings.GetCurrentSettings()
		if cfg.HideOnClose {
			window.Hide()
		} else {
			GetApp().Quit()
		}
	})
}
