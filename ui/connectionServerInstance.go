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

func getServerInstanceComponents() (*widget.Label, *widget.Select) {
	if lblServerInstance == nil || selectServerInstance == nil {
		lblServerInstance = widget.NewLabel(locales.Text("con9"))
		selectServerInstance = widget.NewSelect([]string{""}, func(s string) {
			// TODO
		})
		selectServerInstance.SetSelected("")

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageServerInstance)
	}
	return lblServerInstance, selectServerInstance
}

func updateLanguageServerInstance() {
	lblServerInstance.SetText(locales.Text("con9"))
}

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
