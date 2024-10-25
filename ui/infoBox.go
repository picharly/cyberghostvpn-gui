package ui

import (
	"cyberghostvpn-gui/about"
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/logger"
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/tools"
	"fmt"
	"net"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var infoBox *fyne.Container
var textDefNet *canvas.Text
var textDefStatus *canvas.Text
var textDefVersion *canvas.Text
var textNet *canvas.Text
var textStatus *canvas.Text
var textVersion *canvas.Text

// getInfoBox returns the main information box of the application, which
// contains the application title, the current version of CyberGhost VPN,
// the current status of the VPN (connected or disconnected), and the
// current network IP address of the user. The box is updated every second
// to reflect the current status of the VPN.
func getInfoBox() *fyne.Container {

	if infoBox == nil {
		// App Title
		textTitle := canvas.NewText(fmt.Sprintf("%s - v%s", about.AppName, about.AppVersion), resources.ColorYellow)
		textTitle.TextStyle.Bold = true

		// Version
		textDefVersion = canvas.NewText(locales.Text("inf0"), resources.ColorWhite)
		version := cg.GetVersion()
		if tools.StringContainsNumber(version) {
			textVersion = canvas.NewText(version, resources.ColorYellow)
		} else {
			textVersion = canvas.NewText(locales.Text("inf7"), resources.ColorRed)
		}
		versionBox := layout.NewHBoxLayout()
		versionContainer := container.New(versionBox, textDefVersion, textVersion)

		// Status
		textDefStatus = canvas.NewText(locales.Text("inf1"), resources.ColorWhite)
		textStatus = canvas.NewText(locales.Text("inf7"), resources.ColorWhite)
		statusBox := layout.NewHBoxLayout()
		statusContainer := container.New(statusBox, textDefStatus, textStatus)

		// Network
		textDefNet = canvas.NewText(locales.Text("inf8"), resources.ColorWhite)
		textNet = canvas.NewText("", resources.ColorWhite)
		netBox := layout.NewHBoxLayout()
		netContainer := container.New(netBox, textDefNet, textNet)
		updateNetwork()

		// Update Status
		updateStatus()

		// Create Status Box
		infoBox = container.NewHBox(getCyberGhostLogo(), layout.NewSpacer(), container.NewVBox(layout.NewSpacer(), textTitle, versionContainer, statusContainer, netContainer, layout.NewSpacer()))
	}

	// Enable refresh
	go refresh()

	// Add update method to current trigger
	locales.GetTrigger().AddMethod(updatelanguageInfoBox)

	return infoBox
}

// refresh is a goroutine that periodically updates the current network
// IP address and the current status of the VPN (connected or disconnected).
// It is called by the getInfoBox function and runs until the application is
// terminated.
func refresh() {
	for {
		updateNetwork()
		updateStatus()

		time.Sleep(time.Second * 1)
	}
}

// updatelanguageInfoBox is a function that updates the labels of the main
// information box when the language is changed.
func updatelanguageInfoBox() {
	textDefNet.Text = locales.Text("inf8")
	textDefStatus.Text = locales.Text("inf1")
	textDefVersion.Text = locales.Text("inf0")
}

// updateNetwork is a function that periodically updates the current network
// IP address.
func updateNetwork() {
	if textNet != nil {
		if ip, err := tools.GetLocalIPAddresses(net.FlagPointToPoint); err != nil {
			logger.Errorf("cannot get local IP: %v", err)
			textNet.Text = locales.Text("inf7")
			textNet.Color = resources.ColorOrange
		} else if len(ip) > 0 {
			textNet.Text = ip[0].String()
			textNet.Color = resources.ColorGreen
			GetApp().SetIcon(resources.GetCyberGhostIcon())
		} else {
			textNet.Text = "-" //locales.Text("inf3")
			textNet.Color = resources.ColorRed
			GetApp().SetIcon(resources.GetCyberGhostIconWhite())
		}

		textNet.Refresh()
	}
}

// updateStatus is a function that periodically updates the current status of
// the VPN (connected, disconnected, connecting, disconnecting, not installed).
func updateStatus() {
	status := cg.GetCurrentState()

	switch status {
	case cg.Connected:
		textStatus.Text = locales.Text("inf2")
		textStatus.Color = resources.ColorGreen
	case cg.Disconnected:
		textStatus.Text = locales.Text("inf3")
		textStatus.Color = resources.ColorRed
	case cg.Connecting:
		textStatus.Text = locales.Text("inf4")
		textStatus.Color = resources.ColorOrange
	case cg.Disconnecting:
		textStatus.Text = locales.Text("inf5")
		textStatus.Color = resources.ColorOrange
	case cg.NotInstalled:
		textStatus.Text = locales.Text("inf6")
		textStatus.Color = resources.ColorRed
	default:
		textStatus.Text = locales.Text("inf7")
		textStatus.Color = resources.ColorWhite
	}

	if status == cg.NotInstalled {
		textVersion.Text = locales.Text("inf7")
	} else {
		version := cg.GetVersion()
		if tools.StringContainsNumber(version) {
			textVersion.Text = version
		} else {
			textVersion.Text = locales.Text("inf7")
		}
	}

	textStatus.Refresh()
}
