package resources

import "fyne.io/fyne/v2"

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

func GetWarningIcon() fyne.Resource {
	return fyne.NewStaticResource("warning", WarningIcon)
}
