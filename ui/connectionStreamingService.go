package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"
	"fmt"

	"fyne.io/fyne/v2/widget"
)

var lblStreamingService *widget.Label
var selectStreamingService *widget.Select

// emptyStreamingServiceSelect clears the options of the select widget for the streaming service
// and resets the selected option to an empty string.
func emptyStreamingServiceSelect() {
	selectStreamingService.SetOptions([]string{""})
	selectStreamingService.SetSelected("")
}

// getStreamingServiceComponents returns a Label and a Select widget to select
// the streaming service for CyberGhost. The Select widget is populated with
// the available streaming services for the current country.
// The Select widget is also connected to a function that updates the selected
// streaming service when a new streaming service is selected. The function to
// update the selected streaming service is triggered by the locales trigger.
// The function resets the selected city and server instance when a new streaming
// service is selected and if a country is selected, it updates the cities.
func getStreamingServiceComponents() (*widget.Label, *widget.Select) {
	if lblStreamingService == nil || selectStreamingService == nil {
		lblStreamingService = widget.NewLabel(locales.Text("con10"))
		selectStreamingService = widget.NewSelect([]string{""}, func(s string) {
			if !firstLoad {
				// Reset
				emptyServerInstanceSelect()
				emptyCitySelect()

				// Apply new value
				cg.SetSelectedStreamingService(s)
				if len(s) > 0 {
					if len(selectCountry.Selected) > 0 {
						updateCities(&cg.SelectedCountry)
					}
				}
			}
		})

		// Default
		if cg.SelectedServiceType == string(cg.CG_SERVER_TYPE_STREAMING) {
			updateStreamingServices(false)
		} else {
			selectStreamingService.SetSelected("")
		}

		// Automatic Enable/Disable
		go _automaticEnableDisable(selectStreamingService)

	}
	return lblStreamingService, selectStreamingService
}

// updateLanguageStreamingService updates the label of the select widget of the streaming service component with the new text translated with the current language.
func updateLanguageStreamingService() {
	lblStreamingService.SetText(locales.Text("con10"))
}

// updateStreamingServices updates the select widget of the streaming service component with the available streaming services for the current country.
// The function shows a loading popup while it is updating the streaming services and resets the selected streaming service if the user has selected a new country.
// After updating the select widget, the function resets the selected server instance and city.
func updateStreamingServices(popup bool) {

	// Show loading popup
	if popup {
		showPopupLoading()
		defer removeLoadingWait()
	}

	// Update
	countryCode := cg.SelectedCountry.Code
	if len(loadingStreamingServiceCountry) > 0 {
		countryCode = loadingStreamingServiceCountry
		fmt.Printf("Updating streaming country: %s\n", countryCode)
	}
	services := make([]string, 0)
	services = append(services, "")
	for _, s := range *cg.GetStreamingServices(countryCode) {
		services = append(services, s.Service)
	}
	selectStreamingService.SetOptions(services)

	if cg.SelectedServiceType == string(cg.CG_SERVER_TYPE_STREAMING) {
		selectStreamingService.SetSelected(loadingStreamingService)
	} else {
		selectStreamingService.SetSelected("")
	}
	loadingStreamingService = ""
	loadingStreamingServiceCountry = ""
}
