package main

import (
	"cyberghostvpn-gui/about"
	"cyberghostvpn-gui/logger"
	"cyberghostvpn-gui/ui"
)

func main() {
	// Initiliazing logger
	//logger.AddLoggerUIWriter(ui.GetLogWriter(), cfg.GetTimeFormat())
	logger.LoggerInit(nil)
	if about.DevelopmentMode {
		logger.SetLogLevel("debug")
	} else {
		logger.SetLogLevel("info")
	}

	// Start UI
	ui.GetMainWindow().ShowAndRun()
}
