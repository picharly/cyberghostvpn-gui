package ui

import (
	"cyberghostvpn-gui/locales"

	"fyne.io/fyne/v2/container"
)

var tabs *container.AppTabs
var itemConnection *container.TabItem
var itemSettings *container.TabItem

// getTabs returns the main tabs of the application. The tabs include the
// "Connection" tab and the "Settings" tab. The tabs are updated every time the
// language of the application is changed.
func getTabs() *container.AppTabs {
	if tabs == nil {

		itemConnection = container.NewTabItem(locales.Text("con0"), getConnectionBox())
		itemSettings = container.NewTabItem(locales.Text("set0"), getSettingsBox())

		tabs = container.NewAppTabs(
			itemConnection,
			itemSettings,
		)

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageTabs)
	}

	return tabs
}

// updateLanguageTabs is a method used by the locales trigger to update the
// text of the tabs of the application whenever the language of the application
// is changed.
func updateLanguageTabs() {
	itemConnection.Text = locales.Text("con0")
	itemSettings.Text = locales.Text("set0")
}
