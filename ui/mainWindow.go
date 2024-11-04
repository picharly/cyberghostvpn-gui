package ui

import (
	"cyberghostvpn-gui/about"
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/settings"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

var desktopApp desktop.App
var mainApp fyne.App
var mainWindow fyne.Window

/* Public Functions */

// GetApp returns the main Fyne app instance. If the instance is not yet created, it
// creates a new one with the default dark theme and the CyberGhost VPN icon.
func GetApp() fyne.App {
	if mainApp == nil {
		mainApp = app.NewWithID(about.AppID)
		mainApp.Settings().SetTheme(&resources.DarkTheme{Theme: theme.DefaultTheme()})
		mainApp.SetIcon(resources.GetCyberGhostIcon())
	}
	return mainApp
}

// GetMainWindow returns the main window of the application. If the window is not yet created,
// it creates a new one with the default size and the content set to the main content of the
// application. The window is not resizable and the tray icon is set if the setting is enabled.
func GetMainWindow() fyne.Window {

	if mainWindow == nil {
		// Check for last profile
		cfg, _ := settings.GetCurrentSettings()
		if len(cfg.LastProfile.CountryCode) > 0 {
			// Set value to load
			loadingCountry = cfg.LastProfile.CountryName
			loadingStreamingServiceCountry = cfg.LastProfile.CountryCode
			cg.SetSelectedCountry(cg.GetCountry(cfg.LastProfile.CountryCode))
			loadingStreamingService = cfg.LastProfile.StreamingService
			cg.SetSelectedStreamingService(cfg.LastProfile.StreamingService)
			loadingCity = cfg.LastProfile.City
			loadingServerInstance = cfg.LastProfile.Server
			loadingProtocol = cfg.LastProfile.Protocol
			cg.SetSelectedProtocol(cfg.LastProfile.Protocol)
			loadingServiceType = cfg.LastProfile.ServiceType
			cg.SetSelectedServiceType(cfg.LastProfile.ServiceType)
			loadingVPNService = cfg.LastProfile.VPNService
			cg.SetSelectedVPNService(cfg.LastProfile.VPNService)
		}

		// Create main window
		mainWindow = GetApp().NewWindow(fmt.Sprintf("%s - v%s", about.AppName, about.AppVersion))
		mainWindow.SetFixedSize(true)
		mainWindow.SetContent(getMainContent())

		// Set tray icon
		if cfg.TrayIcon {
			setTrayIcon(mainWindow)
		}

	}

	return mainWindow
}

/* Private Functions */

// getMainContent returns the main content of the application window.
func getMainContent() *fyne.Container {
	vBox := layout.NewVBoxLayout()
	return container.New(vBox, getInfoBox(), getTabs())
}

// setTrayIcon sets the tray icon of the given window. If the window is nil,
// it does nothing. If the window is not nil, it sets the tray icon with a
// menu containing two items: "Hide" and "Show". When the "Hide" item is
// clicked, the window is hidden. When the "Show" item is clicked, the window
// is shown. It also sets the window close intercept to either hide the window
// or quit the application based on the "Hide on close" setting.
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
		desktopApp = desk
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
