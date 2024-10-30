package ui

import (
	"image/color"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type loadingPopup struct {
	widget.BaseWidget
	content fyne.CanvasObject
	canvas  fyne.Canvas
}

var popupLoading *loadingPopup
var mu = sync.Mutex{}
var waitFor int

// addLoadingWait increases the wait for counter for the loading popup.
// This must be called before any action that triggers the loading popup.
// A call to addLoadingWait must be paired with a call to removeLoadingWait.
func addLoadingWait() {
	mu.Lock()
	defer mu.Unlock()
	waitFor++
}

// removeLoadingWait decreases the wait for counter for the loading popup.
// This must be called after any action that triggered the loading popup.
// A call to removeLoadingWait must be paired with a call to addLoadingWait.
func removeLoadingWait() {
	mu.Lock()
	defer mu.Unlock()
	if waitFor > 0 {
		waitFor--
	}
}

// newLoadingPopUp creates a new loading popup.
// The popup is centered on the canvas and is closed automatically when the wait for counter is zero.
// The content of the popup is the provided CanvasObject.
// The popup is only created once, and only the content is updated on subsequent calls.
// A call to showPopupLoading must be paired with a call to addLoadingWait.
func newLoadingPopUp(content fyne.CanvasObject, canvas fyne.Canvas) *loadingPopup {
	popupLoading = &loadingPopup{content: content, canvas: canvas}
	popupLoading.ExtendBaseWidget(popupLoading)
	return popupLoading
}

// Center positions the popup in the center of the canvas.
// It calculates the position so that the popup is centered vertically,
// and the left edge of the popup is at the left edge of the canvas.
func (p *loadingPopup) Center() {
	// Calculate the center position
	popupSize := p.MinSize()
	centerPos := fyne.NewPos(
		0,
		(p.canvas.Size().Height-popupSize.Height)/2,
	)

	// Show the popup at the calculated position
	p.Move(centerPos)
}

// CreateRenderer implements the fyne.Widget interface.
// It creates a renderer that displays the loading popup on top of a black background.
// The background is 80% opaque.
func (p *loadingPopup) CreateRenderer() fyne.WidgetRenderer {
	background := canvas.NewRectangle(color.NRGBA{R: 0, G: 0, B: 0, A: 104}) // 80% opaque black
	return widget.NewSimpleRenderer(container.NewStack(background, p.content))
}

// Show adds the loading popup to the canvas overlays, making it visible.
func (p *loadingPopup) Show() {
	p.canvas.Overlays().Add(p)
}

// Hide removes the loading popup from the canvas overlays, making it invisible.
func (p *loadingPopup) Hide() {
	p.canvas.Overlays().Remove(p)
}

// showPopupLoading shows a loading popup on the main window.
// The popup is displayed on top of all other windows and shows an animated GIF.
// The popup is centered on the screen and has a black background with 80% opacity.
// The size of the popup is the width of the main window and 90px in height.
// The popup is displayed until the function removeLoadingWait is called, which removes it.
// The function can be called multiple times without showing multiple popups.
func showPopupLoading() {

	// Add a new task to the wait group
	addLoadingWait()

	// Only one popup at a time
	if popupLoading != nil {
		return
	}

	// Content
	gif := getLoadingAnimatedGIF()
	popupContainer := container.NewStack(gif)

	// Build popup
	popup := newLoadingPopUp(popupContainer, GetMainWindow().Canvas())
	popup.Resize(fyne.NewSize(GetMainWindow().Canvas().Size().Width, 90))
	popup.Center()
	popup.Show()

	// Start animation
	gif.Start()

	// Wait for end of loading
	go func(p *loadingPopup) {
		for waitFor > 0 {
			// time.Sleep(time.Millisecond * 50)
		}
		p.Hide()
		popupLoading = nil
	}(popup)

}
