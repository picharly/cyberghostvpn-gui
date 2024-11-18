package ui

import (
	"cyberghostvpn-gui/locales"
	"cyberghostvpn-gui/security"
	"cyberghostvpn-gui/settings"
	"cyberghostvpn-gui/tools"
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Encrypted password (stored in memory only - you must type it once at each run)
var password string

// ShowPopupSudo shows a popup window to enter a sudo password to run a command.
// The command is given as a slice of strings, similar to the arguments passed to os/exec.Command.
// The popup shows a label with the text from the locales, a password entry and two buttons: "OK" and "Cancel".
// When the user clicks on "OK" or presses enter, the command is run with the password entered and the popup is closed.
// If the command fails, an error popup is shown after a short delay.
func ShowPopupSudo(args ...string) {

	// New popup
	var p *widget.PopUp
	cfg, _ := settings.GetCurrentSettings()

	// Action
	action := func(args []string, pwd string) {
		// Show loading popup
		p.Hide()
		showPopupLoading()
		defer removeLoadingWait()
		output, err := tools.RunCommand(args, true, true, pwd)
		if err != nil {
			go func(o string, e error) {
				time.Sleep(time.Millisecond * 25)
				showPopupError(fmt.Errorf("%v\n%s", e, o))
			}(strings.Join(output, "\n"), err)
		}
	}

	// Text
	lblText := widget.NewLabel(locales.Text("sud0"))
	// Password
	inputPwd := widget.NewPasswordEntry()
	inputPwd.OnSubmitted = func(v string) {
		if cfg.KeepPassMem {
			password, _ = security.Encrypt(v)
		} else {
			password = ""
		}
		action(args, v)
	}
	if len(password) > 0 {
		decrypt, _ := security.Decrypt(password)
		inputPwd.SetText(decrypt)
		go func(p string, a ...string) {
			time.Sleep(time.Millisecond * 250)
			action(a, p)
		}(decrypt, args...)
	}
	// Checkbox
	checkbox := widget.NewCheck(locales.Text("sud1"), func(b bool) {
		cfg.KeepPassMem = b
		settings.WriteCurrentSettings()
	})
	checkbox.SetChecked(cfg.KeepPassMem)
	// Buttons
	btnOk := widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {
		action(args, inputPwd.Text)
	})

	btnCancel := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		p.Hide()
	})

	// Container
	buttonsContainer := container.NewHBox(btnOk, btnCancel)
	popupContainer := container.NewVBox(lblText, inputPwd, checkbox, layout.NewSpacer(), buttonsContainer)

	// Build popup
	p = widget.NewModalPopUp(popupContainer, GetMainWindow().Canvas())
	p.Resize(fyne.NewSize(300, 90))
	p.Show()
}
