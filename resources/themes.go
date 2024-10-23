package resources

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type DarkTheme struct {
	fyne.Theme
}

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
