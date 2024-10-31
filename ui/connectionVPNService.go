package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"

	"fyne.io/fyne/v2/widget"
)

var lblService *widget.Label
var selectService *widget.Select

// getConnectionVPNServiceComponents returns the label and select widget for the VPN service selection.
// The widget is created only once and the default value is set to the first option.
// The language update method is added to the current trigger.
func getConnectionVPNServiceComponents() (*widget.Label, *widget.Select) {
	if lblService == nil || selectService == nil {

		// Selection of VPN service
		lblService = widget.NewLabel(locales.Text("con4"))
		services := make([]string, 0)
		for k, _ := range cg.VPNServiceOptions {
			services = append(services, k)
		}
		selectService = widget.NewSelect(services, func(s string) {
			if !firstLoad {
				cg.SetSelectedVPNService(s)
			}
		})

		// Default Value
		defaultValue := string(cg.CG_SERVICE_TYPE_OPENVPN)
		if len(loadingVPNService) > 0 {
			defaultValue = loadingVPNService
		}
		loadingVPNService = ""
		selectService.SetSelected(defaultValue)
		cg.SetSelectedVPNService(defaultValue)

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageVPNService)
	}
	return lblService, selectService
}

// updateLanguageVPNService updates the language for the VPN service selection label.
func updateLanguageVPNService() {
	lblService.SetText(locales.Text("con4"))
}
