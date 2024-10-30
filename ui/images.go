package ui

import (
	"cyberghostvpn-gui/resources"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/x/fyne/widget"
)

// getCyberGhostLogo returns a logo of CyberGhost VPN as a canvas.Image.
// The returned image is created from a static resource and has its original size.
func getCyberGhostLogo() *canvas.Image {
	image := canvas.NewImageFromResource(fyne.NewStaticResource("cyberghostvpn_logo.png", resources.CyberGhostLogoPng))
	image.FillMode = canvas.ImageFillOriginal
	return image
}

// getLoadingAnimatedGIF returns an animated GIF canvas object to be used in loading popups.
// The returned GIF is loaded from a static resource and has its size set to 64x64.
func getLoadingAnimatedGIF() *widget.AnimatedGif {
	gif, err := widget.NewAnimatedGifFromResource(resources.GetLoadingGIF())
	if err != nil {
		return nil
	}
	gif.Resize(fyne.NewSize(64, 64))
	return gif
}
