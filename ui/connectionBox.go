package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/resources"
	"cyberghostvpn-gui/settings"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var connectionBox *fyne.Container
var firstLoad = true

var selectedCountry resources.Country
var selectedCity resources.City
var selectedServer resources.Server

func getConnectionBox() *fyne.Container {
	if connectionBox == nil {

		// Profiles
		var btnDelProfile *widget.Button
		var btnAddProfile *widget.Button
		var selectService *widget.Select
		var selectServer *widget.Select
		var selectConnection *widget.Select
		var selectCity *widget.Select
		var selectCountry *widget.Select
		var selectServerInstance *widget.Select

		lblProfile := widget.NewLabel(locales.Text("con1"))
		profiles := make([]string, 0)
		for _, p := range *settings.GetProfiles() {
			profiles = append(profiles, p.Name)
		}
		btnDelProfile = widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {

		})
		selectProfile := widget.NewSelectEntry(profiles)
		selectProfile.Resize(fyne.NewSize(400, 0))
		selectProfile.SetPlaceHolder(locales.Text("con2"))
		selectProfile.OnSubmitted = func(s string) {
			fmt.Printf("Profile selected: %s\n", s)
			// TODO
		}
		btnDelProfile.Disable()
		btnAddProfile = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
			selectProfile.SetOptions(append(profiles, locales.Text("con3"))) //Append(locales.Text("con3"))
		})
		selectProfile.OnChanged = func(s string) {
			if len(s) > 0 && strings.Compare(s, locales.Text("con2")) != 0 {
				btnDelProfile.Enable()
			} else {
				btnDelProfile.Disable()
			}
		}
		containerProfiles := container.NewHBox(selectProfile, btnAddProfile, btnDelProfile)

		// Connection Form

		// Selection of VPN service
		lblService := widget.NewLabel(locales.Text("con4"))
		selectService = widget.NewSelect([]string{string(cg.CG_SERVICE_TYPE_OPENVPN), string(cg.CG_SERVICE_TYPE_WIREGUARD)}, func(s string) {
			// TODO
		})
		selectService.SetSelected(string(cg.CG_SERVICE_TYPE_OPENVPN))

		// Selection of Server Type
		lblServer := widget.NewLabel(locales.Text("con5"))
		selectServer = widget.NewSelect([]string{string(cg.CG_SERVER_TYPE_TRAFFIC), string(cg.CG_SERVER_TYPE_STREAMING), string(cg.CG_SERVER_TYPE_TORRENT)}, func(s string) {
			if !firstLoad {
				selectServerInstance.SetOptions([]string{""})
				selectServerInstance.SetSelected("")
				selectCity.SetOptions([]string{""})
				selectCity.SetSelected("")
				countries := make([]string, 0)
				countries = append(countries, "")
				sel := cg.CG_SERVER_TYPE_TRAFFIC
				switch s {
				case string(cg.CG_SERVER_TYPE_TRAFFIC):
					sel = cg.CG_SERVER_TYPE_TRAFFIC
				case string(cg.CG_SERVER_TYPE_STREAMING):
					sel = cg.CG_SERVER_TYPE_STREAMING
				case string(cg.CG_SERVER_TYPE_TORRENT):
					sel = cg.CG_SERVER_TYPE_TORRENT
				}
				for _, c := range *cg.GetCountries(sel) {
					countries = append(countries, c.Name)
				}
				selectCountry.SetOptions(countries)
				selectCountry.Selected = ""
			}
		})
		selectServer.SetSelected(string(cg.CG_SERVER_TYPE_TRAFFIC))

		// Selection of Connection Protocol
		lblConnection := widget.NewLabel(locales.Text("con6"))
		selectConnection = widget.NewSelect([]string{string(cg.CG_CONNECTION_UDP), string(cg.CG_CONNECTION_TCP)}, func(s string) {
			// TODO
		})
		selectConnection.SetSelected(string(cg.CG_CONNECTION_UDP))

		// Selection of Country
		lblCountry := widget.NewLabel(locales.Text("con7"))
		countries := make([]string, 0)
		countries = append(countries, "")
		for _, c := range *cg.GetCountries(cg.CG_SERVER_TYPE_TRAFFIC) {
			countries = append(countries, c.Name)
		}
		selectCountry = widget.NewSelect(countries, func(s string) {
			if !firstLoad {
				// Reset
				selectServerInstance.SetOptions([]string{""})
				selectServerInstance.SetSelected("")
				selectCity.SetSelected("")
				selectCity.Disable()

				if len(s) > 0 {
					selectedCountry = cg.GetCountry(s)
					if len(selectedCountry.Name) > 0 {
						go func(sel *widget.Select, selCountry *resources.Country) {
							cities := make([]string, 0)
							cities = append(cities, "")
							for _, c := range *cg.GetCities(cg.CG_SERVER_TYPE_TRAFFIC, selCountry.Code) {
								cities = append(cities, c.Name)
							}
							sel.SetOptions(cities)
							sel.SetSelected("")
							sel.Enable()
						}(selectCity, &selectedCountry)
					}
				}
			}
		})
		selectCountry.SetSelected("")

		// Selection of City
		lblCity := widget.NewLabel(locales.Text("con8"))
		selectCity = widget.NewSelect([]string{""}, func(s string) {
			if !firstLoad {
				// Reset
				selectServerInstance.SetOptions([]string{""})
				selectServerInstance.SetSelected("")
				selectServerInstance.Disable()

				if len(s) > 0 {
					selectedCity = cg.GetCity(s)
					if len(selectedCity.Name) > 0 {
						go func(sel *widget.Select, selCountry *resources.Country, selCity *resources.City) {
							srv := make([]string, 0)
							srv = append(srv, "")
							for _, c := range *cg.GetServers(cg.CG_SERVER_TYPE_TRAFFIC, selCountry.Code, selCity.Name) {
								srv = append(srv, fmt.Sprintf("%s (%s)", c.Instance, c.Load))
							}
							sel.SetOptions(srv)
							sel.SetSelected("")
							sel.Enable()
						}(selectServerInstance, &selectedCountry, &selectedCity)
					}
				}
			}
		})
		selectCity.SetSelected("")

		// Selection of Server Instance
		lblServerInstance := widget.NewLabel(locales.Text("con9"))
		selectServerInstance = widget.NewSelect([]string{""}, func(s string) {
			// TODO
		})
		selectServerInstance.SetSelected("")

		// CG_OTHER_CONNECTION   cgCommand = "--connection"   // needs a connection type (UDP, TCP)
		// CG_OTHER_COUNTRY_CODE cgCommand = "--country-code" // needs a country code
		// CG_OTHER_CITY         cgCommand = "--city"         // needs a city name
		// CG_OTHER_SERVER       cgCommand = "--server"       // needs a server name

		form := container.New(layout.NewFormLayout(), lblProfile, containerProfiles, lblService, selectService, lblServer, selectServer, lblConnection, selectConnection, lblCountry, selectCountry, lblCity, selectCity, lblServerInstance, selectServerInstance)

		connectionBox = container.NewVBox(form)

		firstLoad = false
	}

	return connectionBox
}
