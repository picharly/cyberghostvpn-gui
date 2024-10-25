package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"

	"fyne.io/fyne/v2/widget"
)

var lblCountry *widget.Label
var selectCountry *widget.Select

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
				selectServerInstance.SetOptions([]string{""})
				selectServerInstance.SetSelected("")
				selectCity.SetSelected("")
				selectCity.Disable()

				if len(s) > 0 {
					cg.SetSelectedCountry(cg.GetCountry(s))
					if len(cg.SelectedCountry.Name) > 0 {
						go updateCities(&cg.SelectedCountry)
					}
				}
			}
		})
		selectCountry.SetSelected("")

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageCountry)
	}
	return lblCountry, selectCountry
}

func updateCountries(serverType cg.CgServerType) {
	countries := make([]string, 0)
	countries = append(countries, "")
	for _, c := range *cg.GetCountries(serverType) {
		countries = append(countries, c.Name)
	}
	selectCountry.SetOptions(countries)
	selectCountry.Selected = ""
	selectCountry.Enable()
}

func updateLanguageCountry() {
	lblCountry.SetText(locales.Text("con7"))
}
