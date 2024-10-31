package resources

import (
	"embed"
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
