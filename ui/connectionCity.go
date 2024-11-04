package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/resources"

	"fyne.io/fyne/v2/widget"
)

var lblCity *widget.Label
var selectCity *widget.Select

// emptyCitySelect empties the city select widget and selects the first option (empty string).
func emptyCitySelect() {
	selectCity.SetOptions([]string{""})
	selectCity.SetSelected("")
}

// getCityComponents returns a Label and a Select widget to select a city for CyberGhost.
// The Select widget is populated with the available cities for the current country.
// The Select widget is also connected to a function that updates the selected city when a new city is selected.
// The function to update the selected city is triggered by the locales trigger.
func getCityComponents() (*widget.Label, *widget.Select) {
	if lblCity == nil || selectCity == nil {
		lblCity = widget.NewLabel(locales.Text("con8"))
		selectCity = widget.NewSelect([]string{""}, func(s string) {
			if !firstLoad {
				// Reset
				emptyServerInstanceSelect()

				cg.SetSelectedCity(cg.GetCity(s))
				if len(s) > 0 {
					if len(cg.SelectedCity.Name) > 0 {
						updateServerInstances(&cg.SelectedCountry, &cg.SelectedCity)
					}
				}
			}
		})

		// Default
		if len(loadingCity) > 0 {
			updateCities(&cg.SelectedCountry)
		} else {
			selectCity.SetSelected("")
		}

		// Automatic Enable/Disable
		go _automaticEnableDisable(selectCity)

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageCity)
	}
	return lblCity, selectCity
}

// updateCities updates the city select with the available cities for the current country.
// It loads the cities if they are not already loaded and updates the selected city if a city was previously selected.
// It also shows a loading popup while it is updating the cities.
func updateCities(selCountry *resources.Country) {

	// Show loading popup
	showPopupLoading()
	defer removeLoadingWait()

	// Update
	cities := make([]string, 0)
	cities = append(cities, "")
	for _, c := range *cg.GetCities(cg.CgServerType(selectServerType.Selected), selCountry.Code) {
		cities = append(cities, c.Name)
	}
	selectCity.SetOptions(cities)
	if len(loadingCity) > 0 {
		selectCity.SetSelected(loadingCity)
	} else {
		selectCity.SetSelected("")
	}
	loadingCity = ""
}

// updateLanguageCity updates the label of the city select with the current language.
func updateLanguageCity() {
	lblCity.SetText(locales.Text("con8"))
}
