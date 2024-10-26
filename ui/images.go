package ui

import (
	"cyberghostvpn-gui/resources"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

// getCyberGhostLogo returns a logo of CyberGhost VPN as a canvas.Image.
// The returned image is created from a static resource and has its original size.
func getCyberGhostLogo() *canvas.Image {
	image := canvas.NewImageFromResource(fyne.NewStaticResource("cyberghostvpn_logo.png", resources.CyberGhostLogoPng))
	image.FillMode = canvas.ImageFillOriginal
	return image
}
