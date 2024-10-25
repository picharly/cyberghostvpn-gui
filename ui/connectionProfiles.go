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

var containerProfiles *fyne.Container
var lblProfile *widget.Label
var selectProfile *widget.SelectEntry

// getConnectionProfilesComponents returns a label and a container that contains a select widget for selecting a profile,
// a button to add a new profile and a button to delete the selected profile.
// The returned label is the title of the select widget.
// The select widget is loaded with the names of all the profiles that are defined in the settings.
func getConnectionProfilesComponents() (*widget.Label, *fyne.Container) {
	if lblProfile == nil || selectProfile == nil || containerProfiles == nil {

		var btnDelProfile *widget.Button
		var btnAddProfile *widget.Button

		lblProfile = widget.NewLabel(locales.Text("con1"))
		profiles := make([]string, 0)
		for _, p := range *settings.GetProfiles() {
			profiles = append(profiles, p.Name)
		}
		btnDelProfile = widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {

		})
		selectProfile = widget.NewSelectEntry(profiles)
		selectProfile.Resize(fyne.NewSize(400, 0))
		selectProfile.SetPlaceHolder(locales.Text("con2"))
		selectProfile.OnSubmitted = func(s string) {
			// TODO
		}
		btnDelProfile.Disable()
		btnAddProfile = widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
			selectProfile.SetOptions(append(profiles, locales.Text("con3")))
		})
		selectProfile.OnChanged = func(s string) {
			if len(s) > 0 && strings.Compare(s, locales.Text("con2")) != 0 {
				btnDelProfile.Enable()
			} else {
				btnDelProfile.Disable()
			}
		}
		containerProfiles = container.NewHBox(selectProfile, btnAddProfile, btnDelProfile)

		// Add update method to current trigger
		locales.GetTrigger().AddMethod(updateLanguageProfiles)
	}
	return lblProfile, containerProfiles
}

// updateLanguageProfiles updates the label of the select widget of the profiles component with the new text translated with the current language.
func updateLanguageProfiles() {
	lblProfile = widget.NewLabel(locales.Text("con1"))
}
