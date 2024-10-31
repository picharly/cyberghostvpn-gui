package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"

	"fyne.io/fyne/v2/widget"
)

var lblServerType *widget.Label
var selectServerType *widget.Select

// getServerTypeComponents returns a Label and a Select widget to select
// the server type for CyberGhost. The Select widget is populated with
// the three server types: traffic, streaming, torrent. The Select widget
// is also connected to a function that updates the countries when a
// server type is changed. The function to update the countries is
// triggered by the locales trigger.
func getServerTypeComponents() (*widget.Label, *widget.Select) {
	if lblServerType == nil || selectServerType == nil {
		lblServerType = widget.NewLabel(locales.Text("con5"))
		serverTypes := make([]string, 0)
		for k, _ := range cg.ServerTypeOptions {
			serverTypes = append(serverTypes, k)
		}
		selectServerType = widget.NewSelect(serverTypes, func(s string) {
			if !firstLoad {

				// Reset current selection
				emptyServerInstanceSelect()
				emptyCitySelect()
				emptyStreamingServiceSelect()

				// Get new selection
				cg.SetSelectedServiceType(s)

				// Update countries
				if len(selectCountry.Selected) < 1 {
					updateCountries(cg.GetServerType(s))
				} else {
					if s == string(cg.CG_SERVER_TYPE_STREAMING) {
						updateStreamingServices()
					} else {
						emptyStreamingServiceSelect()
						updateCities(&cg.SelectedCountry)
					}
				}

			}
		})

		// Default option
		defaultOption := cg.CG_SERVER_TYPE_TRAFFIC
		if len(loadingServiceType) > 0 {
			defaultOption = cg.GetServerType(loadingServiceType)
		}
		loadingServiceType = ""
		selectServerType.SetSelected(string(defaultOption))
		cg.SetSelectedServiceType(string(defaultOption))

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageServerType)
	}
	return lblServerType, selectServerType
}

// updateLanguageServerType updates the label of the server type select with the current language
func updateLanguageServerType() {
	lblServerType.SetText(locales.Text("con5"))
}
