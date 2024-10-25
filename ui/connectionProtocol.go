package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"

	"fyne.io/fyne/v2/widget"
)

var lblConnection *widget.Label
var selectConnection *widget.Select

func getConnectionProtocolComponents() (*widget.Label, *widget.Select) {
	if lblConnection == nil || selectConnection == nil {
		lblConnection = widget.NewLabel(locales.Text("con6"))
		selectConnection = widget.NewSelect([]string{string(cg.CG_CONNECTION_UDP), string(cg.CG_CONNECTION_TCP)}, func(s string) {
			// TODO
		})
		selectConnection.SetSelected(string(cg.CG_CONNECTION_UDP))

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageConnectionProtocol)
	}
	return lblConnection, selectConnection
}

func updateLanguageConnectionProtocol() {
	lblConnection.SetText(locales.Text("con6"))
}
