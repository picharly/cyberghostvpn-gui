package cg

import (
	"cyberghostvpn-gui/tools"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var outputLabel *widget.Entry

func TestSudo() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Root Command Executor")

	outputLabel = widget.NewMultiLineEntry()
	outputLabel.SetPlaceHolder("Command output will appear here...")

	runButton := widget.NewButton("Run 'dmesg' with gksudo", func() {
		output, err := tools.RunCommandWithGksudo("dmesg")
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}
		outputLabel.SetText(output)
		outputLabel.Refresh()
	})

	content := container.NewVBox(
		widget.NewLabel("Execute a command with gksudo:"),
		runButton,
		outputLabel,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
