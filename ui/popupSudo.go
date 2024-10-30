package ui

import (
	"fmt"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func showPopupSudo() {
	password := widget.NewPasswordEntry()
	dialog.ShowCustomConfirm("Enter sudo password", "OK", "Cancel", password, func(ok bool) {
		if ok {
			// Use the password with sudo
			cmd := exec.Command("sudo", "-S", "")
			cmd.Stdin = strings.NewReader(password.Text + "\n")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Command failed: %v", err)
			} else {
				fmt.Printf("Success %s", string(output))
			}
		}
	}, GetMainWindow())
}
