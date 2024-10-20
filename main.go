package main

import (
	"cyberghostvpn-gui/about"
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/logger"
	"cyberghostvpn-gui/settings"
	"cyberghostvpn-gui/ui"
)

// main is the entry point of the application. It initializes settings, loads
// the appropriate locale, configures the logging system, and starts the
// graphical user interface. It sets the logging level based on whether the
// application is in development mode. Additionally, the function contains
// commented-out code for testing local IP addresses.
func main() {
	// Read settings & load locale
	if cfg, err := settings.GetCurrentSettings(); err == nil && len(cfg.Language) > 0 {
		locales.Init(cfg.Language)
	} else {
		locales.Init("")
	}

	// Initiliaze logger
	//logger.AddLoggerUIWriter(ui.GetLogWriter(), cfg.GetTimeFormat())
	logger.LoggerInit(nil)
	if about.DevelopmentMode {
		logger.SetLogLevel("debug")
	} else {
		logger.SetLogLevel("info")
	}

	// test ip
	// ips, err := tools.GetLocalIPAddresses(net.FlagPointToPoint)
	// if err == nil && len(ips) > 0 {
	// 	for _, ip := range ips {
	// 		fmt.Printf("IP: %s\n", ip.String())
	// 	}
	// }
	// time.Sleep(time.Minute)

	// Start UI
	ui.GetMainWindow().ShowAndRun()
}
