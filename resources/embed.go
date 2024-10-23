package resources

import (
	"embed"

	"fyne.io/fyne/v2"
)

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

func GetCyberGhostIcon() fyne.Resource {
	return fyne.NewStaticResource("cyberghostvpn_icon", CyberGhostIconOriginal)
}
func GetCyberGhostIconError() fyne.Resource {
	return fyne.NewStaticResource("cyberghostvpn_icon", CyberGhostIconRed)
}

func GetCyberGhostIconWhite() fyne.Resource {
	return fyne.NewStaticResource("cyberghostvpn_icon", CyberGhostIconWhite)
}
