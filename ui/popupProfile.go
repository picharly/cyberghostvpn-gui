package ui

import (
	"cyberghostvpn-gui/cg"
	"cyberghostvpn-gui/locales"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

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
