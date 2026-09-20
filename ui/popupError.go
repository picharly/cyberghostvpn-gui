package ui

import (
	"fyne.io/fyne/v2/dialog"
)

// showPopupError shows an error popup with the given error message.
// The popup is displayed in the main window and the error message
// is displayed as the content of the popup.
// It must be called from the Fyne main thread (use fyne.Do from goroutines).
func showPopupError(err error) {

	// Error dialog
	d := dialog.NewError(err, GetMainWindow())
	d.Show()
}
