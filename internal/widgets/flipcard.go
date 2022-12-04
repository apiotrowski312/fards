package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type flipCard struct {
	widget.BaseWidget

	background *canvas.Rectangle
	label      *widget.Label

	frontLabel, backLabel string
	flipped               bool
}

func NewFlipCard(front, back string) *flipCard {
	w := &flipCard{
		frontLabel: front,
		backLabel:  back,
	}
	w.ExtendBaseWidget(w)
	return w
}

func (w *flipCard) Reset(front, back string) {
	w.flipped = false
	w.frontLabel = front
	w.backLabel = back

	w.label.SetText(w.frontLabel)
}

func (w *flipCard) Tapped(_ *fyne.PointEvent) {
	if w.flipped {
		w.label.SetText(w.frontLabel)
	} else {
		w.label.SetText(w.backLabel)
	}
	w.flipped = !w.flipped
}

func (w *flipCard) TappedSecondary(_ *fyne.PointEvent) {}

func (w *flipCard) CreateRenderer() fyne.WidgetRenderer {
	w.background = canvas.NewRectangle(theme.ButtonColor())
	w.label = widget.NewLabel(w.frontLabel)

	w.label.Alignment = fyne.TextAlignCenter
	w.label.Wrapping = fyne.TextWrapWord

	flipCardContainer := container.NewMax(w.background, container.NewVBox(layout.NewSpacer(), w.label, layout.NewSpacer()))
	return widget.NewSimpleRenderer(flipCardContainer)
}
