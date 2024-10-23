package ui

import (
	"cyberghostvpn-gui/locales"

	"fyne.io/fyne/v2/container"
)

var tabs *container.AppTabs

func getTabs() *container.AppTabs {
	if tabs == nil {

		tabs = container.NewAppTabs(
			container.NewTabItem(locales.Text("con0"), getConnectionBox()),
			container.NewTabItem(locales.Text("set0"), getSettingsBox()),
		)
	}

	return tabs
}
