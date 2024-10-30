package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/logger"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// showPopupLoadingProfile shows a popup window to load a profile.
// The popup shows the name of the selected profile and a progress bar.
// The progress bar is updated as the profile is loaded.
// The popup is closed after the profile is loaded or if loading takes more than 6 seconds.
func showPopupLoadingProfile() {
	// New popup
	var popup *widget.PopUp

	// Read current name
	name := selectProfile.Selected
	// Text
	lblCountry := widget.NewLabel(locales.Text("pro6", locales.Variable{Name: "Name", Value: loadingCountry}))
	lblCity := widget.NewLabel(locales.Text("pro7", locales.Variable{Name: "Name", Value: loadingCity}))
	lblCity.Hide()
	lblServerInstance := widget.NewLabel(locales.Text("pro8", locales.Variable{Name: "Name", Value: loadingServerInstance}))
	lblServerInstance.Hide()

	// Progress bar
	prg := widget.NewProgressBar()

	// Container
	popupContainer := container.NewVBox(lblCountry, lblCity, lblServerInstance, layout.NewSpacer(), prg)

	// Build popup
	popup = widget.NewModalPopUp(popupContainer, GetMainWindow().Canvas())
	popup.Resize(fyne.NewSize(300, 90))
	popup.Show()

	// Thread update
	go func(lc *widget.Label, lsi *widget.Label, pr *widget.ProgressBar, p *widget.PopUp) {
		start := time.Now()
		for {
			if loadingCountry == "" {
				lc.Show()
				prg.SetValue(33.3)
			}
			if loadingCity == "" {
				lc.Show()
				lsi.Show()
				prg.SetValue(66.6)
			}
			if loadingServerInstance == "" {
				lc.Show()
				lsi.Show()
				prg.SetValue(100.0)
				break
			}
			if time.Since(start) > 6*time.Second {
				logger.Errorf("timeout while loading profile %v", name)
				break
			}
			popup.Refresh()
		}
		p.Hide()

	}(lblCity, lblServerInstance, prg, popup)
}

// showPopupProfileDelete shows a popup window to confirm the deletion of the selected profile.
// The popup shows the name of the profile and two buttons: "OK" and "Cancel".
// If the user clicks on "OK", the profile is deleted and the popup is closed.
// If the user clicks on "Cancel", the popup is closed without deleting the profile.
func showPopupProfileDelete() {

	// New popup
	var p *widget.PopUp

	// Read current name
	name := selectProfile.Selected
	// Text
	lblText := widget.NewLabel(locales.Text("pro5", locales.Variable{Name: "Name", Value: name}))

	// Buttons
	btnOk := widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {
		cg.DeleteProfile(name)
		p.Hide()
		updateProfiles()
	})

	btnCancel := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		p.Hide()
	})

	// Container
	buttonsContainer := container.NewHBox(btnOk, btnCancel)
	popupContainer := container.NewVBox(lblText, layout.NewSpacer(), buttonsContainer)

	// Build popup
	p = widget.NewModalPopUp(popupContainer, GetMainWindow().Canvas())
	p.Resize(fyne.NewSize(300, 90))
	p.Show()
}

// showPopupProfileSave shows a popup window to save a profile.
// The popup shows an input field with the current name of the selected profile.
// The user can enter a new name and click on "Save" to save the profile.
// If the user clicks on "Cancel", the popup is closed without saving the profile.
// If the user enters a name with less than one character, the input field is marked as invalid.
func showPopupProfileSave() {

	// New popup
	var p *widget.PopUp

	// Form
	previousName := ""
	lblName := widget.NewLabel(locales.Text("pro1"))
	input := widget.NewEntry()
	input.SetPlaceHolder(locales.Text("pro3"))
	// Read current name
	if len(selectProfile.Selected) > 0 {
		input.SetText(selectProfile.Selected)
		previousName = selectProfile.Selected
	}
	// Validate input
	input.Validator = func(s string) error {
		if len(s) < 1 {
			return fmt.Errorf(locales.Text("pro2"))
		}
		return nil
	}

	// Buttons
	btnOk := widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {
		if err := input.Validate(); err != nil {
			input.FocusLost()
			return
		} else {
			cg.SaveProfile(input.Text, previousName)
			p.Hide()
			updateProfiles()
		}
	})

	btnCancel := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		p.Hide()
	})

	// Container
	formContainer := container.New(
		layout.NewFormLayout(),
		lblName, input)
	buttonsContainer := container.NewHBox(btnOk, btnCancel)
	popupContainer := container.NewVBox(formContainer, layout.NewSpacer(), buttonsContainer)

	// Build popup
	p = widget.NewModalPopUp(popupContainer, GetMainWindow().Canvas())
	p.Resize(fyne.NewSize(300, 90))
	input.FocusGained()
	p.Canvas.Focus(input)
	p.Show()
}
