package ui

import (
	"fyne.io/fyne/v2/dialog"
)

func showPopupError(err error) {

	// Error dialog
	d := dialog.NewError(err, GetMainWindow())
	d.Show()
}
