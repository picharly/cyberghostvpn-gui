package resources

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type DarkTheme struct {
	fyne.Theme
}

// Color returns the color associated with the given theme color name and variant.
//
// Note: CyberGhost VPN Dark theme overrides the default colors as follows:
// - Background: ColorBlack
// - Foreground: ColorWhite
// - Primary: ColorYellow
func (m DarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return ColorBlack
	case theme.ColorNameForeground:
		return ColorWhite
	case theme.ColorNamePrimary:
		return ColorYellow
	}
	return theme.DefaultTheme().Color(name, variant)
}
