package resources

import "image/color"

type CustomColor color.Color

var ColorBlue CustomColor = color.RGBA{R: 0, G: 0, B: 255, A: 255}
var ColorBlack CustomColor = color.RGBA{R: 22, G: 22, B: 23, A: 255}
var ColorGreen CustomColor = color.RGBA{R: 0, G: 255, B: 0, A: 255}
var ColorOrange CustomColor = color.RGBA{R: 255, G: 128, B: 0, A: 255}
var ColorRed CustomColor = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var ColorYellow CustomColor = color.RGBA{R: 254, G: 203, B: 0, A: 255}
var ColorWhite CustomColor = color.RGBA{R: 194, G: 194, B: 194, A: 255}
