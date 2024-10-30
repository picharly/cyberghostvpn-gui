package resources

import (
	"embed"

	"fyne.io/fyne/v2"
)

// Animations
//
//go:embed animations/loading.gif
var LoadingGIF []byte

// Pictures
//
//go:embed pictures/cyberghostvpn_logo.png
var CyberGhostLogoPng []byte

// Flags
// Source: https://github.com/lipis/flag-icons
//
//go:embed flags/*.svg
var srcFlags embed.FS

// Icons
//
//go:embed icons/cyberghostvpn_icon_original.png
var CyberGhostIconOriginal []byte

//go:embed icons/cyberghostvpn_icon_red.png
var CyberGhostIconRed []byte

//go:embed icons/cyberghostvpn_icon_transparent.png
var CyberGhostIconTransparent []byte

//go:embed icons/cyberghostvpn_icon_white.png
var CyberGhostIconWhite []byte

// GetCyberGhostIcon returns a Fyne resource representing the original CyberGhost VPN icon.
// The icon is embedded as a static resource and can be used throughout the application.
func GetCyberGhostIcon() fyne.Resource {
	return fyne.NewStaticResource("cyberghostvpn_icon", CyberGhostIconOriginal)
}

// GetCyberGhostIconError returns a Fyne resource representing the CyberGhost VPN icon in red color.
// This icon is used when an error occurs in the application.
func GetCyberGhostIconError() fyne.Resource {
	return fyne.NewStaticResource("cyberghostvpn_icon", CyberGhostIconRed)
}

// GetCyberGhostIconWhite returns a Fyne resource representing the CyberGhost VPN icon in white color.
// This icon can be used as a visual element within the application where a neutral color is needed.
func GetCyberGhostIconWhite() fyne.Resource {
	return fyne.NewStaticResource("cyberghostvpn_icon", CyberGhostIconWhite)
}

// GetLoadingGIF returns a Fyne resource representing the animated GIF displayed in a loading popup.
// The GIF is embedded as a static resource and can be used throughout the application.
func GetLoadingGIF() fyne.Resource {
	return fyne.NewStaticResource("loading", LoadingGIF)
}
