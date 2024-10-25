package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/resources"
	"fmt"

	"fyne.io/fyne/v2/widget"
)

var lblServerInstance *widget.Label
var selectServerInstance *widget.Select

// getServerInstanceComponents returns a Label and a Select widget to select
// the server instance for CyberGhost. The Select widget is populated with
// the available server instances for the current country and city.
// The Select widget is also connected to a function that updates the selected
// server when a new server instance is selected. The function to update the
// selected server is triggered by the locales trigger.
func getServerInstanceComponents() (*widget.Label, *widget.Select) {
	if lblServerInstance == nil || selectServerInstance == nil {
		lblServerInstance = widget.NewLabel(locales.Text("con9"))
		selectServerInstance = widget.NewSelect([]string{""}, func(s string) {
			if !firstLoad {
				if len(s) > 0 {
					cg.SetSelectedServer(cg.GetServer(s))
				}
			}
		})

		// Default
		selectServerInstance.SetSelected("")
		if firstLoad {
			selectServerInstance.Disable()
		}

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageServerInstance)
	}
	return lblServerInstance, selectServerInstance
}

// updateLanguageServerInstance updates the label of the server instance select with the current language
func updateLanguageServerInstance() {
	lblServerInstance.SetText(locales.Text("con9"))
}

// updateServerInstances updates the server instance select with the available server instances for the current country and city.
func updateServerInstances(selCountry *resources.Country, selCity *resources.City) {
	srv := make([]string, 0)
	srv = append(srv, "")
	for _, c := range *cg.GetServers(cg.CG_SERVER_TYPE_TRAFFIC, selCountry.Code, selCity.Name) {
		srv = append(srv, fmt.Sprintf("%s (%s)", c.Instance, c.Load))
	}
	selectServerInstance.SetOptions(srv)
	selectServerInstance.SetSelected("")
	selectServerInstance.Enable()
}
