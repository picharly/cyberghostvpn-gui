package ui

import (
	"cyberghostvpn-gui/about"
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/settings"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
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

		// Set tray icon
		cfg, _ := settings.GetCurrentSettings()
		if cfg.TrayIcon {
			setTrayIcon(mainWindow)
		}

		// Hide on tray
		if cfg.StartTray {
			go func(window fyne.Window) {
				time.Sleep(time.Millisecond * 10)
				mainWindow.Hide()
			}(mainWindow)
		}
	}

	return mainWindow
}

/* Private Functions */

func getApp() fyne.App {
	if mainApp == nil {
		mainApp = app.New()
		mainApp.SetIcon(resources.GetCyberGhostIcon())
	}
	return mainApp
}

func getMainContent() *fyne.Container {
	text2 := canvas.NewText("2", resources.ColorWhite)
	text3 := canvas.NewText("3", resources.ColorWhite)
	vBox := layout.NewVBoxLayout()
	return container.New(vBox, getInfoBox(), text2, text3)
}

func setTrayIcon(window fyne.Window) {
	if window == nil {
		return
	}
	if desk, ok := getApp().(desktop.App); ok {
		m := fyne.NewMenu(about.AppName,
			fyne.NewMenuItem("Hide", func() { window.Hide() }),
			fyne.NewMenuItem("Show", func() { window.Show() }),
		)
		desk.SetSystemTrayMenu(m)
	}

	cfg, _ := settings.GetCurrentSettings()
	if cfg.HideOnTray {
		window.SetCloseIntercept(func() {
			window.Hide()
		})
	}
}
