package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"

	"fyne.io/fyne/v2/widget"
)

var lblStreamingService *widget.Label
var selectStreamingService *widget.Select

func getStreamingServiceComponents() (*widget.Label, *widget.Select) {
	if lblStreamingService == nil || selectStreamingService == nil {
		lblStreamingService = widget.NewLabel(locales.Text("con10"))
		selectStreamingService = widget.NewSelect([]string{""}, func(s string) {
			if !firstLoad {
				if len(s) > 0 {
					cg.SetSelectedStreamingService(s)
				}

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

		// Default
		selectStreamingService.SetSelected("")
		selectStreamingService.Disable()

	}
	return lblStreamingService, selectStreamingService
}

func updateLanguageStreamingService() {
	lblStreamingService.SetText(locales.Text("con10"))
}

func updateStreamingServices() {
	services := make([]string, 0)
	services = append(services, "")
	for _, s := range *cg.GetStreamingServices(cg.SelectedCountry.Code) {
		services = append(services, s.Service)
	}
	selectStreamingService.SetOptions(services)
	selectStreamingService.SetSelected("")
}
