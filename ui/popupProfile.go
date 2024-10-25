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

func showPopupProfileSave() {

	// New popup
	var p *widget.PopUp

	// Form
	lblName := widget.NewLabel(locales.Text("pro1"))
	input := widget.NewEntry()
	// input.Resize(fyne.NewSize(300, 30))
	input.SetPlaceHolder(locales.Text("pro3"))
	if len(selectProfile.Selected) > 0 {
		input.SetText(selectProfile.Selected)
	}
	input.Validator = func(s string) error {
		if len(s) < 1 {
			return fmt.Errorf(locales.Text("pro2"))
		}
		return nil
	}

	// Buttons
	btnOk := widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {
		input.FocusLost()
		if err := input.Validate(); err != nil {
			return
		} else {
			cg.SaveProfile(input.Text)
			p.Hide()
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
	// popupContainer.Resize(fyne.NewSize(400, 200))

	// Build popup
	p = widget.NewModalPopUp(popupContainer, GetMainWindow().Canvas())
	p.Resize(fyne.NewSize(300, 90))
	input.FocusGained()
	p.Show()
}
