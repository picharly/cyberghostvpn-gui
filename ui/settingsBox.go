package ui

import (
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/settings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var settingsBox *fyne.Container

func getSettingsBox() *fyne.Container {

	if settingsBox == nil {

		cfg, _ := settings.GetCurrentSettings()

		// Language
		lblLanguage := widget.NewLabel(locales.Text("set1"))
		languages := make([]string, 0)
		for _, l := range locales.GetLocales() {
			languages = append(languages, l.Name)
		}
		comboboxLanguage := widget.NewSelect(languages, func(s string) {
			for _, l := range locales.GetLocales() {
				if s == l.Name {
					locales.Init(l.Code)
					cfg, _ := settings.GetCurrentSettings()
					cfg.Language = l.Code
					settings.WriteCurrentSettings()
					break
				}
			}
		})
		comboboxLanguage.SetSelected(locales.GetCurrentLocale().Name)

		// Enable tray icon
		lblTrayIcon := widget.NewLabel(locales.Text("set2"))
		checkboxTrayIcon := widget.NewCheck("", func(b bool) {
			cfg, _ := settings.GetCurrentSettings()
			cfg.TrayIcon = b
			settings.WriteCurrentSettings()
		})
		checkboxTrayIcon.SetChecked(cfg.TrayIcon)

		// Start on tray
		lblStartOnTray := widget.NewLabel(locales.Text("set3"))
		checkboxStartOnTray := widget.NewCheck("", func(b bool) {
			cfg, _ := settings.GetCurrentSettings()
			cfg.HideOnStart = b
			settings.WriteCurrentSettings()
		})
		checkboxStartOnTray.SetChecked(cfg.HideOnStart)

		// Hide on tray
		lblHideOnTray := widget.NewLabel(locales.Text("set4"))
		checkboxHideOnTray := widget.NewCheck("", func(b bool) {
			cfg, _ := settings.GetCurrentSettings()
			cfg.HideOnClose = b
			settings.WriteCurrentSettings()
		})
		checkboxHideOnTray.SetChecked(cfg.HideOnClose)

		form := container.New(
			layout.NewFormLayout(),
			lblLanguage,
			comboboxLanguage,
			lblTrayIcon,
			checkboxTrayIcon,
			lblStartOnTray,
			checkboxStartOnTray,
			lblHideOnTray,
			checkboxHideOnTray,
		)
		settingsBox = container.NewHBox(form)
	}

	// 	HideOnTray  bool                `json:"hide_on_tray"`
	// Language    string              `json:"language"`
	// LastProfile []Profile           `json:"last_profile"`
	// StartTray   bool                `json:"start_tray"`
	// TrayIcon    bool                `json:"tray_icon"`

	return settingsBox
}
