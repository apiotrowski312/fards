package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type editor struct {
	router router
	canvas fyne.Canvas

	cards cardsStore
}

func NewEditor(cards cardsStore, router router, canvas fyne.Canvas) *editor {
	return &editor{
		router: router,
		canvas: canvas,

		cards: cards,
	}
}

func (e editor) Layout() fyne.CanvasObject {
	e.getToLastCard()

	return e.container()
}

func (e editor) container() fyne.CanvasObject {
	labelCardsStat := widget.NewLabel("")
	labelCardsStat.Alignment = fyne.TextAlignTrailing

	questionEntry := widget.NewEntry()
	answerEntry := widget.NewEntry()

	questionEntry.SetPlaceHolder("Question")
	answerEntry.SetPlaceHolder("Answer")

	questionEntry.OnSubmitted = func(s string) {
		e.canvas.Focus(answerEntry)
	}

	prevBtn := widget.NewButton("Prev", nil)
	nextBtn := widget.NewButton("Next", nil)

	currentCard, _ := e.cards.GetStats()
	if currentCard == 1 {
		prevBtn.Disable()
	}

	setAll := func() {
		q, a := e.cards.GetCard()
		questionEntry.SetText(q)
		answerEntry.SetText(a)

		current, max := e.cards.GetStats()
		labelCardsStat.SetText(fmt.Sprintf("%v/%v", current, max))

		if q == "" {
			e.canvas.Focus(questionEntry)
		}
	}

	nextFunc := func() {
		e.cards.Update(questionEntry.Text, answerEntry.Text)

		isOk := e.cards.Next()
		if !isOk {
			e.cards.Add("", "")
			e.cards.Next()
		}

		prevBtn.Enable()

		setAll()
	}

	prevBtn.OnTapped = func() {
		e.cards.Update(questionEntry.Text, answerEntry.Text)

		isOk := e.cards.Previous()
		if !isOk {
			prevBtn.Disable()
			return
		}
		currentCard, _ := e.cards.GetStats()
		if currentCard == 1 {
			prevBtn.Disable()
		}

		setAll()
	}

	answerEntry.OnSubmitted = func(s string) {
		nextFunc()
	}

	nextBtn.OnTapped = func() {
		nextFunc()
	}

	setAll()

	back := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		e.cards.Update(questionEntry.Text, answerEntry.Text)
		e.router.Pop()
	})

	removeBtn := widget.NewButton("Remove card", func() {
		e.cards.RemoveCurrentCard()
		setAll()
	})

	top := container.NewBorder(nil, nil, back, removeBtn, labelCardsStat)
	bottom := container.NewGridWithColumns(2, prevBtn, nextBtn)
	center := container.NewGridWithRows(2, questionEntry, answerEntry)
	return container.NewBorder(top, bottom, nil, nil, center)
}

func (e editor) getToLastCard() {
	for e.cards.Next() {
	}
}
