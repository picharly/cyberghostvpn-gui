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

// Components with text
var checkboxConnectStartup *widget.Check
var checkboxHideOnTray *widget.Check
var checkboxLoadLastProfile *widget.Check
var checkboxStartOnTray *widget.Check
var checkboxStopVPN *widget.Check
var checkboxTrayIcon *widget.Check

var lblConnectStartup *widget.Label
var lblHideOnTray *widget.Label
var lblLanguage *widget.Label
var lblLoadLastProfile *widget.Label
var lblStartOnTray *widget.Label
var lblStopVPN *widget.Label
var lblTrayIcon *widget.Label

// getSettingsBox returns the settings box of the application, which contains
// the settings UI components.  The first time it is called, the UI components
// are created and the settingsBox variable is set.  On subsequent calls, the
// value of settingsBox is returned immediately.
//
// The settings box is a horizontal box containing the following components: the
// language text and form, the tray icon text and form, the start on tray text
// and form, the hide on tray text and form, and the stop VPN on exit text and
// form.
func getSettingsBox() *fyne.Container {

	if settingsBox == nil {
		building := true

		cfg, _ := settings.GetCurrentSettings()

		// Language
		lblLanguage = widget.NewLabel(locales.Text("set1"))
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
					if !building {
						locales.GetTrigger().Activate()
					}
					break
				}
			}
		})
		comboboxLanguage.SetSelected(locales.GetCurrentLocale().Name)

		// Enable tray icon
		lblTrayIcon = widget.NewLabel(locales.Text("set2"))
		checkboxTrayIcon = widget.NewCheck("", func(b bool) {
			cfg, _ := settings.GetCurrentSettings()
			cfg.TrayIcon = b

			// Update other components
			if checkboxStartOnTray != nil {
				if checkboxTrayIcon.Checked {
					checkboxHideOnTray.Enable()
					checkboxStartOnTray.Enable()
				} else {
					checkboxHideOnTray.Disable()
					checkboxHideOnTray.SetChecked(false)
					checkboxStartOnTray.Disable()
					checkboxStartOnTray.SetChecked(false)
				}
			}

			settings.WriteCurrentSettings()
		})
		checkboxTrayIcon.SetChecked(cfg.TrayIcon)

		// Start on tray
		lblStartOnTray = widget.NewLabel(locales.Text("set3"))
		checkboxStartOnTray = widget.NewCheck("", func(b bool) {
			cfg, _ := settings.GetCurrentSettings()
			cfg.HideOnStart = b
			settings.WriteCurrentSettings()
		})
		checkboxStartOnTray.SetChecked(cfg.HideOnStart)
		if !cfg.TrayIcon {
			checkboxStartOnTray.Disable()
		}

		// Hide on tray
		lblHideOnTray = widget.NewLabel(locales.Text("set4"))
		checkboxHideOnTray = widget.NewCheck("", func(b bool) {
			cfg, _ := settings.GetCurrentSettings()
			cfg.HideOnClose = b
			settings.WriteCurrentSettings()
		})
		checkboxHideOnTray.SetChecked(cfg.HideOnClose)
		if !cfg.TrayIcon {
			checkboxHideOnTray.Disable()
		}

		// Stop VPN on exit
		lblStopVPN = widget.NewLabel(locales.Text("set5"))
		checkboxStopVPN = widget.NewCheck("", func(b bool) {
			cfg, _ := settings.GetCurrentSettings()
			cfg.StopVPNOnExit = b
			settings.WriteCurrentSettings()
		})
		checkboxStopVPN.SetChecked(cfg.StopVPNOnExit)

		// Load last profile
		lblLoadLastProfile = widget.NewLabel(locales.Text("set6"))
		checkboxLoadLastProfile = widget.NewCheck("", func(b bool) {
			cfg, _ := settings.GetCurrentSettings()
			cfg.LoadLastProfile = b

			// Update other components
			if checkboxConnectStartup != nil {
				if checkboxLoadLastProfile.Checked {
					checkboxConnectStartup.Enable()
				} else {
					checkboxConnectStartup.Disable()
					checkboxConnectStartup.SetChecked(false)
				}
			}

			settings.WriteCurrentSettings()
		})
		checkboxLoadLastProfile.SetChecked(cfg.LoadLastProfile)

		// Connect on startup
		lblConnectStartup = widget.NewLabel(locales.Text("set7"))
		checkboxConnectStartup = widget.NewCheck("", func(b bool) {
			cfg, _ := settings.GetCurrentSettings()
			cfg.ConnectStartup = b
			settings.WriteCurrentSettings()
		})
		checkboxConnectStartup.SetChecked(cfg.ConnectStartup)
		if !cfg.LoadLastProfile {
			checkboxConnectStartup.Disable()
		}

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
			// lblLoadLastProfile,
			// checkboxLoadLastProfile,
			lblConnectStartup,
			checkboxConnectStartup,
			lblStopVPN,
			checkboxStopVPN,
		)
		settingsBox = container.NewHBox(form)

		building = false

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageSettings)
	}

	return settingsBox
}

// updateLanguageSettings updates the language-related texts in the settings box
// when the language is changed.
func updateLanguageSettings() {
	// Language
	lblLanguage.SetText(locales.Text("set1"))
	lblTrayIcon.SetText(locales.Text("set2"))
	lblStartOnTray.SetText(locales.Text("set3"))
	lblHideOnTray.SetText(locales.Text("set4"))
	lblStopVPN.SetText(locales.Text("set5"))
	lblLoadLastProfile.SetText(locales.Text("set6"))
	lblConnectStartup.SetText(locales.Text("set7"))
}
