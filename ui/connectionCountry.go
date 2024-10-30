package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"

	"fyne.io/fyne/v2/widget"
)

var lblCountry *widget.Label
var selectCountry *widget.Select

// emptyCountrySelect empties the country select widget and sets the selected item to an empty string.
func emptyCountrySelect() {
	selectCountry.SetOptions([]string{""})
	selectCountry.Selected = ""
}

// getCountryComponents returns a label and a select widget to select the country for CyberGhost.
// The select widget is populated with the countries that are available for the traffic server type.
// The select widget is also connected to a function that updates the selected country when a new country is selected.
// The function to update the selected country is triggered by the locales trigger.
// The function also updates the cities and streaming services when a new country is selected.
func getCountryComponents() (*widget.Label, *widget.Select) {
	if lblCountry == nil || selectCountry == nil {
		lblCountry = widget.NewLabel(locales.Text("con7"))
		countries := make([]string, 0)
		countries = append(countries, "")
		for _, c := range *cg.GetCountries(cg.CG_SERVER_TYPE_TRAFFIC) {
			countries = append(countries, c.Name)
		}
		selectCountry = widget.NewSelect(countries, func(s string) {
			if !firstLoad {
				// Reset
				emptyServerInstanceSelect()
				emptyCitySelect()
				emptyStreamingServiceSelect()

				// Set selection
				cg.SetSelectedCountry(cg.GetCountry(s))

				// Show flag
				if len(cg.SelectedCountry.Name) > 0 {
					setFlag(cg.SelectedCountry.Code)
				} else {
					setFlag("")
				}

				// Refresh
				if len(s) > 0 {
					// Streaming service
					if selectServerType.Selected == string(cg.CG_SERVER_TYPE_STREAMING) {
						updateStreamingServices()
						selectStreamingService.Enable()
					} else {
						selectStreamingService.SetSelected("")
						selectStreamingService.Disable()

						// City
						if len(cg.SelectedCountry.Name) > 0 {
							updateCities(&cg.SelectedCountry)
						}
					}
				} else {
					btnConnect.Disable()
				}

			} else {
				selectCountry.Disable()
			}
		})

		// Automatic Enable/Disable
		go _automaticEnableDisable(selectCountry)

		if len(loadingCountry) > 0 {
			selectCountry.SetSelected(loadingCountry)
		} else {
			selectCountry.SetSelected("")
		}
		loadingCountry = ""

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageCountry)
	}
	return lblCountry, selectCountry
}

// updateCountries updates the country select widget with the list of countries
// available for the given server type. It displays a loading popup while the
// countries are being retrieved. Once updated, the select widget is enabled,
// and the list of countries is set as options with no selection by default.
//
// Parameters:
//
//	serverType: The server type for which to retrieve the list of countries.
func updateCountries(serverType cg.CgServerType) {

	// Show loading popup
	showPopupLoading()
	defer removeLoadingWait()

	// Update
	countries := make([]string, 0)
	countries = append(countries, "")
	for _, c := range *cg.GetCountries(serverType) {
		countries = append(countries, c.Name)
	}
	selectCountry.SetOptions(countries)
	selectCountry.Selected = ""
	selectCountry.Enable()
}

// updateLanguageCountry updates the label of the country select widget
// with the text translated into the current language.
func updateLanguageCountry() {
	lblCountry.SetText(locales.Text("con7"))
}
