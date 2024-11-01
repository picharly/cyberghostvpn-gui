package ui

import (
	"cyberghostvpn-gui/about"
	"cyberghostvpn-gui/locales"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ShowRequirementsPopup creates a new window displaying a list of missing dependencies. The window is centered on screen and has a fixed size. The window is closed when the user clicks on the "OK" button and the application exits with a non-zero return code.
func ShowRequirementsPopup(missing []string) {
	// Create a new window
	a := GetApp()
	w := a.NewWindow(fmt.Sprintf("%s - v%s", about.AppName, about.AppVersion))
	w.SetFixedSize(true)
	w.SetOnClosed(func() {
		a.Quit()
		os.Exit(1)
	})
	w.CenterOnScreen()

	// Create a new content
	content := []fyne.CanvasObject{}
	content = append(content, widget.NewLabel(locales.Text("req1")))
	txtMissing := ""
	for _, v := range missing {
		txtMissing += "- " + v + "\n"
	}
	content = append(content, widget.NewLabel(txtMissing))
	content = append(content, widget.NewButton(locales.Text("gen7"), func() {
		w.Close()
	}))

	// Set the content
	w.SetContent(container.NewHBox(getWarningPicture(), container.NewVBox(content...)))

	// Show the window
	w.ShowAndRun()
}
