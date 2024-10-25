package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

var connectionBox *fyne.Container
var firstLoad = true

// getConnectionBox generates the connection box UI component.
//
// The first time it is called, the UI components are created and the
// connectionBox variable is set.  On subsequent calls, the value of
// connectionBox is returned immediately.
//
// The connection box is a vertical box containing the following
// components: the profile text and form, the VPN service text and form,
// the server type text and form, the connection protocol text and form,
// the country text and form, the city text and form, and the server
// instance text and form.
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
