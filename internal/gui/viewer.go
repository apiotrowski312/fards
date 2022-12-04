package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/apiotrowski312/fards/internal/widgets"
)

type viewer struct {
	router router
	canvas fyne.Canvas

	cards cardsStore
}

func NewViewer(cards cardsStore, router router, canvas fyne.Canvas) *viewer {
	cards.Shuffle()
	return &viewer{
		router: router,
		canvas: canvas,

		cards: cards,
	}
}

func (v viewer) Layout() fyne.CanvasObject {
	return container.NewBorder(nil, v.bottom(), nil, nil, v.center())
}

func (v viewer) center() fyne.CanvasObject {
	label := widget.NewLabel("")
	label.Alignment = fyne.TextAlignTrailing

	flipCard := widgets.NewFlipCard("", "")
	center := container.NewGridWithRows(3, label, flipCard, layout.NewSpacer())

	nextBtn := widget.NewButton("Next", nil)
	nextFlipCard := func(next bool) {
		if next && !v.cards.Next() {
			v.reload()
		}

		current, max := v.cards.GetStats()
		if current == max {
			nextBtn.SetText("Start over")
		}

		front, back := v.cards.GetCard()
		flipCard.Reset(front, back)
		label.SetText(fmt.Sprintf("%v/%v", current, max))
	}

	nextBtn.OnTapped = func() {
		nextFlipCard(true)
	}

	nextFlipCard(false)
	return container.NewBorder(nil, nextBtn, nil, nil, center)
}

func (v viewer) bottom() *fyne.Container {
	left := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		v.router.Pop()
	})
	right := widget.NewButton("Shuffle", func() {
		v.cards.Shuffle()
		v.reload()
	})

	return container.NewBorder(nil, nil, left, right, container.NewMax())
}

func (v viewer) reload() {
	for v.cards.Previous() {
	}
	v.router.Reload()
}
