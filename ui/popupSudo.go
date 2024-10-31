package ui

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func showPopupSudo(args ...string) {
	password := widget.NewPasswordEntry()
	dialog.ShowCustomConfirm("Enter sudo password", "OK", "Cancel", password, func(ok bool) {
		var procErr error
		if ok {
			// Use the password with sudo
			cmd := exec.Command("sudo", "-S", "--", "/usr/sbin/cyberghostvpn", "--connect")
			cmd.Stdin = strings.NewReader(password.Text)
			output, err := cmd.CombinedOutput()
			if err != nil {
				procErr = fmt.Errorf("%v (%s)", err, string(output))
			}
		}
		if procErr != nil {
			time.Sleep(time.Millisecond * 25) // Wait for popup to close
			showPopupError(procErr)
		}
	}, GetMainWindow())
}
