package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"

	"fyne.io/fyne/v2/widget"
)

var lblServerType *widget.Label
var selectServerType *widget.Select

func getServerTypeComponents() (*widget.Label, *widget.Select) {
	if lblServerType == nil || selectServerType == nil {
		lblServerType = widget.NewLabel(locales.Text("con5"))
		selectServerType = widget.NewSelect([]string{string(cg.CG_SERVER_TYPE_TRAFFIC), string(cg.CG_SERVER_TYPE_STREAMING), string(cg.CG_SERVER_TYPE_TORRENT)}, func(s string) {
			if !firstLoad {
				selectServerInstance.SetOptions([]string{""})
				selectServerInstance.SetSelected("")
				selectCity.SetOptions([]string{""})
				selectCity.SetSelected("")
				countries := make([]string, 0)
				countries = append(countries, "")
				sel := cg.CG_SERVER_TYPE_TRAFFIC
				switch s {
				case string(cg.CG_SERVER_TYPE_TRAFFIC):
					sel = cg.CG_SERVER_TYPE_TRAFFIC
				case string(cg.CG_SERVER_TYPE_STREAMING):
					sel = cg.CG_SERVER_TYPE_STREAMING
				case string(cg.CG_SERVER_TYPE_TORRENT):
					sel = cg.CG_SERVER_TYPE_TORRENT
				}
				for _, c := range *cg.GetCountries(sel) {
					countries = append(countries, c.Name)
				}
				selectCountry.SetOptions(countries)
				selectCountry.Selected = ""
			}
		})
		selectServerType.SetSelected(string(cg.CG_SERVER_TYPE_TRAFFIC))
		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageServerType)
	}
	return lblServerType, selectServerType
}

func updateLanguageServerType() {
	lblServerType.SetText(locales.Text("con5"))
}
