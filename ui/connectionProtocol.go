package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"

	"fyne.io/fyne/v2/widget"
)

var lblConnection *widget.Label
var selectConnection *widget.Select

// getConnectionProtocolComponents returns a label and a select widget for selecting the connection protocol.
// The returned label is the title of the select widget.
// The select widget is loaded with the names of the connection protocols that are defined in the settings.
// The currently selected protocol is stored in the selectedProtocol variable.
// The method for updating the label with the new text translated with the current language is added to the current trigger.
func getConnectionProtocolComponents() (*widget.Label, *widget.Select) {
	if lblConnection == nil || selectConnection == nil {
		lblConnection = widget.NewLabel(locales.Text("con6"))
		selectConnection = widget.NewSelect([]string{string(cg.CG_CONNECTION_UDP), string(cg.CG_CONNECTION_TCP)}, func(s string) {
			cg.SetSelectedProtocol(s)
		})

		// Default Value
		defaultValue := string(cg.CG_CONNECTION_UDP)
		selectConnection.SetSelected(defaultValue)
		cg.SetSelectedProtocol(defaultValue)

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageConnectionProtocol)
	}
	return lblConnection, selectConnection
}

// updateLanguageConnectionProtocol updates the label of the select widget of the connection protocol component with the new text translated with the current language.
func updateLanguageConnectionProtocol() {
	lblConnection.SetText(locales.Text("con6"))
}
