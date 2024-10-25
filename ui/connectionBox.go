package ui

import (
	"cyberghostvpn-gui/resources"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var connectionBox *fyne.Container
var firstLoad = true

var selectedCountry resources.Country
var selectedCity resources.City
var selectedServer resources.Server

func getConnectionBox() *fyne.Container {
	if connectionBox == nil {

		// Profiles
		profileText, profileForm := getConnectionProfilesComponents()

		// Selection of VPN service
		vpnServiceText, vpnServiceForm := getConnectionVPNServiceComponents()

		// Selection of Server Type
		serverTypeText, serverTypeForm := getServerTypeComponents()

		// Selection of Connection Protocol
		protocolText, protocolForm := getConnectionProtocolComponents()

		// Selection of Country
		countryText, countryForm := getCountryComponents()

		// Selection of City
		cityText, cityForm := getCityComponents()

		// Selection of Server Instance
		serverInstanceText, serverInstanceForm := getServerInstanceComponents()

		// Form
		form := container.New(layout.NewFormLayout(),
			profileText, profileForm,
			vpnServiceText, vpnServiceForm,
			serverTypeText, serverTypeForm,
			protocolText, protocolForm,
			countryText, countryForm,
			cityText, cityForm,
			serverInstanceText, serverInstanceForm)

		connectionBox = container.NewVBox(form)

		firstLoad = false
	}

	return connectionBox
}
