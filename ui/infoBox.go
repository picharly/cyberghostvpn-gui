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

var imgFlag *canvas.Image
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

		// Country
		imgFlag = canvas.NewImageFromResource(fyne.NewStaticResource("flag.png", resources.GetFlag("FR")))
		imgW := 64
		imgH := 48
		imgFlag.SetMinSize(fyne.NewSize(float32(imgW), float32(imgH)))
		imgContainer := container.New(layout.NewCenterLayout(), imgFlag)
		imgContainer.Resize(fyne.NewSize(float32(imgW), float32(imgH)))
		imgFlag.Translucency = 100

		// Update Status
		updateStatus()

		// Create Status Box
		infoBox = container.NewHBox(
			getCyberGhostLogo(),
			layout.NewSpacer(),
			container.NewVBox(
				layout.NewSpacer(),
				textTitle,
				versionContainer,
				statusContainer,
				netContainer,
				imgContainer,
				layout.NewSpacer(),
			),
		)
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
	lastUpdate := time.Now()
	for {
		// Refresh CgVPN stateevery second
		if time.Since(lastUpdate) > time.Millisecond*1000 {
			cg.GetCurrentState()
			lastUpdate = time.Now()
		}

		// Update InfoBox
		updateNetwork()
		updateStatus()
		updateConnectButtonStatus()

		time.Sleep(time.Millisecond * 100)
	}
}

// setFlag sets the flag of the country with the given code in the main information box of the application.
// If the given code is empty, the flag is hidden by setting its transparency to 100.
// Otherwise, the flag is updated with the given code and its transparency is set to 0.
func setFlag(countryCode string) {
	if len(countryCode) < 1 {
		imgFlag.Translucency = 100
	} else {
		imgFlag.Resource = fyne.NewStaticResource(countryCode+".svg", resources.GetFlag(countryCode))
		imgFlag.Translucency = 0
	}
	imgFlag.Refresh()
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
			logger.Errorf("%s %v", locales.Text("err.inf1"), err)
			textNet.Text = locales.Text("inf7")
			textNet.Color = resources.ColorOrange
		} else if len(ip) > 0 {
			textNet.Text = ip[0].String()
			textNet.Color = resources.ColorGreen
			GetApp().SetIcon(resources.GetCyberGhostIcon())
		} else {
			textNet.Text = "-" //locales.Text("inf3")
			textNet.Color = resources.ColorRed
			GetApp().SetIcon(resources.GetCyberGhostIcon())
		}

		go func() {
			fyne.DoAndWait(func() {
				textNet.Refresh()
			})
		}()
	}
}

// updateStatus is a function that periodically updates the current status of
// the VPN (connected, disconnected, connecting, disconnecting, not installed).
func updateStatus() {
	status := cg.CurrentState

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

	go func() {
		fyne.DoAndWait(func() {
			textStatus.Refresh()
		})
	}()
}
