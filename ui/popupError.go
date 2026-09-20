package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// showPopupError shows an error popup with the given error message.
// The popup is displayed in the main window and the error message
// is displayed as the content of the popup.
// It is safe to call from any goroutine: the dialog is always shown
// on the Fyne main thread.
func showPopupError(err error) {

	// Error dialog
	fyne.Do(func() {
		d := dialog.NewError(err, GetMainWindow())
		d.Show()
	})
}
