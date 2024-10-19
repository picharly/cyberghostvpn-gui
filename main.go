package main

import (
	"cyberghostvpn-gui/about"
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/logger"
	"cyberghostvpn-gui/ui"
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
