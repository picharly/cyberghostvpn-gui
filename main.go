package main

import (
	"cyberghostvpn-gui/about"
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/logger"
	"cyberghostvpn-gui/security"
	"cyberghostvpn-gui/settings"
	"cyberghostvpn-gui/tools"
	"cyberghostvpn-gui/ui"
	"strings"
)

// main is the entry point of the application. It initializes settings, loads
// the appropriate locale, configures the logging system, and starts the
// graphical user interface. It sets the logging level based on whether the
// application is in development mode. Additionally, the function contains
// commented-out code for testing local IP addresses.
func main() {
	// Read settings & load locale
	cfg, err := settings.GetCurrentSettings()
	if err == nil && len(cfg.Language) > 0 {
		locales.Init(cfg.Language)
	} else {
		locales.Init("")
	}

	// Initiliaze logger
	logger.SetDateTimeFormat(locales.Text("date"), locales.Text("time"))
	logger.GetCurrentLogger()
	if about.DevelopmentMode {
		logger.SetLogLevel("debug")
	} else {
		logger.SetLogLevel("info")
	}

	// Load countries
	cg.GetCountries(cg.CG_SERVER_TYPE_TRAFFIC)

	// Check requirements
	if missing, ok := settings.CheckRequirements(); !ok {
		ui.ShowRequirementsPopup(missing)
	}

	// Get current VPN state
	cg.GetCurrentState()

	// Start UI
	if cfg.HideOnStart {
		ui.GetMainWindow().Hide()
		ui.GetApp().Run()
	} else {
		ui.GetMainWindow().ShowAndRun()
	}

	// Disconnect before exit
	//
	//        /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\
	// Because of a limitation of Fyne.io library, it is not possible to catch a Close/Quit event from menu
	// like the on on Tray Icon because it is an automatically generated action.
	// In this case, the app will try to disconnect from VPN but sudo password must has been memorized during
	// this session.
	//        /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\ /!\
	//
	if cfg.StopVPNOnExit {
		args := cg.Disconnect()
		decrypt, _ := security.Decrypt(ui.Password)
		output, err := tools.RunCommand(args, true, true, decrypt)
		if err != nil {
			logger.Errorf("%v\n%s", err, strings.Join(output, "\n"))
		}
	}
}
