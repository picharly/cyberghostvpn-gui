package ui

import (
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/settings"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var btnDelProfile *widget.Button
var btnSaveProfile *widget.Button
var containerProfiles *fyne.Container
var lblProfile *widget.Label
var selectProfile *widget.Select

var loadingCountry string
var loadingCity string
var loadingServerInstance string
var loadingStreamingService string

// getConnectionProfilesComponents returns a label and a container that contains a select widget for selecting a profile,
// a button to add a new profile and a button to delete the selected profile.
// The returned label is the title of the select widget.
// The select widget is loaded with the names of all the profiles that are defined in the settings.
func getConnectionProfilesComponents() (*widget.Label, *fyne.Container) {
	if lblProfile == nil || selectProfile == nil || containerProfiles == nil {

		lblProfile = widget.NewLabel(locales.Text("con1"))
		profiles := make([]string, 0)
		profiles = append(profiles, "")
		for _, p := range *settings.GetProfiles() {
			profiles = append(profiles, p.Name)
		}
		btnDelProfile = widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
			showPopupProfileDelete()
		})
		selectProfile = widget.NewSelect(profiles, func(s string) {
			if len(s) < 1 {
				btnDelProfile.Disable()
			} else {
				btnDelProfile.Enable()
			}
		})
		btnDelProfile.Disable()
		btnSaveProfile = widget.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
			//selectProfile.SetOptions(append(profiles, locales.Text("con3")))
			showPopupProfileSave()
		})
		selectProfile.OnChanged = func(s string) {
			if len(s) > 0 && strings.Compare(s, locales.Text("con2")) != 0 {
				btnDelProfile.Enable()
				p := settings.GetProfile(s)
				if p != nil {
					// Set value to load
					loadingCountry = p.CountryName
					loadingStreamingService = p.StreamingService
					loadingCity = p.City
					loadingServerInstance = p.Server

					// Clear current values
					emptyCountrySelect()

					// Apply profile settings
					selectServerType.SetSelected(p.ServiceType)

					selectService.SetSelected(p.VPNService)
					selectConnection.SetSelected(p.Protocol)
					selectCountry.SetSelected(p.CountryName)
				}
			} else {
				btnDelProfile.Disable()
			}
		}
		containerProfiles = container.NewHBox(selectProfile, btnSaveProfile, btnDelProfile)

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageProfiles)
	}
	return lblProfile, containerProfiles
}

// updateLanguageProfiles updates the label of the select widget of the profiles component with the new text
// translated with the current language.
func updateLanguageProfiles() {
	lblProfile = widget.NewLabel(locales.Text("con1"))
}

// updateProfiles updates the options of the select widget of the profiles component with the names of the profiles
// that are currently in the settings and selects the first option (empty string)
func updateProfiles() {
	currentProfileName := selectProfile.Selected
	profiles := make([]string, 0)
	profiles = append(profiles, "")
	for _, p := range *settings.GetProfiles() {
		profiles = append(profiles, p.Name)
	}
	selectProfile.SetOptions(profiles)
	if len(currentProfileName) > 0 {
		selectProfile.SetSelected(currentProfileName)
	} else {
		selectProfile.SetSelected("")
	}
}
