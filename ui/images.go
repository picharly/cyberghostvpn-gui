package ui

import (
	"cyberghostvpn-gui/resources"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func getCyberGhostLogo() *canvas.Image {
	image := canvas.NewImageFromResource(fyne.NewStaticResource("cyberghostvpn_logo.png", resources.CyberGhostLogoPng))
	image.FillMode = canvas.ImageFillOriginal
	return image
}
