package logic

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// TODO: Let's make history in sync with layouts
// TODO: Improve implementation

type Layoutable interface {
	Layout() fyne.CanvasObject
}

type screenRouter struct {
	current *fyne.Container
	history []fyne.CanvasObject
	layouts []Layoutable

	canvas fyne.Canvas
}

func NewScreenRouter(canvas fyne.Canvas) *screenRouter {
	popup := container.NewCenter()
	return &screenRouter{
		current: container.NewMax(widget.NewLabel("It should not be visible"), popup),

		layouts: []Layoutable{},

		canvas: canvas,
	}
}

func (s *screenRouter) CleanPopup() {
	popup := s.current.Objects[1].(*fyne.Container)
	popup.RemoveAll()
}

func (s *screenRouter) SetPopup(content fyne.CanvasObject) {
	modal := widget.NewModalPopUp(content, s.canvas)
	popup := s.current.Objects[1].(*fyne.Container)
	popup.Add(modal)
}

func (s *screenRouter) Reset(newMain Layoutable) *screenRouter {
	s.current.Objects[0] = newMain.Layout()
	s.history = []fyne.CanvasObject{}
	s.layouts = []Layoutable{newMain}

	return s
}

func (s *screenRouter) Pop() {
	// We cannot pop last screen
	if len(s.history) == 0 {
		return
	}

	lastIndex := len(s.history) - 1

	s.current.Objects[0] = s.history[lastIndex]

	s.history = s.history[:lastIndex]
	s.layouts = s.layouts[:lastIndex+1]

}

func (s *screenRouter) Push(screen Layoutable) {
	s.history = append(s.history, s.current.Objects[0])
	s.layouts = append(s.layouts, screen)

	s.current.Objects[0] = screen.Layout()
}

func (s *screenRouter) Reload() {
	s.current.Objects[0] = s.layouts[len(s.layouts)-1].Layout()
}

func (s *screenRouter) Container() *fyne.Container {
	return s.current
}
