package main

import (
	"cyberghostvpn-gui/about"
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/logger"
)

func main() {
	// Define locale
	locales.Init()

	// Initiliaze logger
	//logger.AddLoggerUIWriter(ui.GetLogWriter(), cfg.GetTimeFormat())
	logger.LoggerInit(nil)
	if about.DevelopmentMode {
		logger.SetLogLevel("debug")
	} else {
		logger.SetLogLevel("info")
	}

	// Test
	// Commande à exécuter avec sudo
	cg.TestSudo()

	//time.Sleep(time.Minute)

	// Start UI
	//ui.GetMainWindow().ShowAndRun()
}
