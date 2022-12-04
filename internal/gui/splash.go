package gui

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/apiotrowski312/fards/assets"
)

const fadeAfter = 1500 * time.Millisecond

// MakeSplash - returns splash screen.
func MakeSplash() fyne.CanvasObject {
	text := canvas.NewText("FARDS", color.White)
	text.TextSize = 50
	text.TextStyle = fyne.TextStyle{Italic: true, Bold: true}

	icon := canvas.NewImageFromResource(assets.ResourceIconWhitePng)
	icon.SetMinSize(fyne.NewSize(100, 100))

	vBox := container.NewVBox(
		container.NewCenter(icon),
		container.NewCenter(text),
	)

	return container.NewMax(
		canvas.NewRectangle(theme.BackgroundColor()),
		container.NewCenter(vBox),
	)
}

// FadeSplash - will remove canvas object.
func FadeSplash(obj fyne.CanvasObject) {
	time.Sleep(fadeAfter)
	obj.Hide()
	obj.Refresh()
}
