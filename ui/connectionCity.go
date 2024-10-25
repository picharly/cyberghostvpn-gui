package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/resources"

	"fyne.io/fyne/v2/widget"
)

var lblCity *widget.Label
var selectCity *widget.Select

func getCityComponents() (*widget.Label, *widget.Select) {
	if lblCity == nil || selectCity == nil {
		lblCity = widget.NewLabel(locales.Text("con8"))
		selectCity = widget.NewSelect([]string{""}, func(s string) {
			if !firstLoad {
				// Reset
				selectServerInstance.SetOptions([]string{""})
				selectServerInstance.SetSelected("")
				selectServerInstance.Disable()

				if len(s) > 0 {
					cg.SetSelectedCity(cg.GetCity(s))
					if len(cg.SelectedCity.Name) > 0 {
						go updateServerInstances(&cg.SelectedCountry, &cg.SelectedCity)
					}
				}
			}
		})

		// Default
		selectCity.SetSelected("")
		if firstLoad {
			selectCity.Disable()
		}

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageCity)
	}
	return lblCity, selectCity
}

func updateCities(selCountry *resources.Country) {
	cities := make([]string, 0)
	cities = append(cities, "")
	for _, c := range *cg.GetCities(cg.CG_SERVER_TYPE_TRAFFIC, selCountry.Code) {
		cities = append(cities, c.Name)
	}
	selectCity.SetOptions(cities)
	selectCity.SetSelected("")
	selectCity.Enable()
}

func updateLanguageCity() {
	lblCity.SetText(locales.Text("con8"))
}
