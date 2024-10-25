package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"

	"fyne.io/fyne/v2/widget"
)

var lblService *widget.Label
var selectService *widget.Select

func getConnectionVPNServiceComponents() (*widget.Label, *widget.Select) {
	if lblService == nil || selectService == nil {

		// Selection of VPN service
		lblService = widget.NewLabel(locales.Text("con4"))
		selectService = widget.NewSelect([]string{string(cg.CG_SERVICE_TYPE_OPENVPN), string(cg.CG_SERVICE_TYPE_WIREGUARD)}, func(s string) {
			// TODO
		})
		selectService.SetSelected(string(cg.CG_SERVICE_TYPE_OPENVPN))

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageVPNService)
	}
	return lblService, selectService
}

// updateLanguageVPNService updates the language for the VPN service selection label.
func updateLanguageVPNService() {
	lblService.SetText(locales.Text("con4"))
}
