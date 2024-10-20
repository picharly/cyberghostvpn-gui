package resources

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed cyberghostvpn_logo.png
var CyberGhostLogoPng []byte

//go:embed cyberghostvpn_icon_white.png
var CyberGhostIconWhite []byte

func GetCyberGhostIcon() fyne.Resource {
	return fyne.NewStaticResource("cyberghostvpn_icon", CyberGhostIconWhite)
}
