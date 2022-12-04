package logic_test

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/apiotrowski312/fards/internal/logic"
	"github.com/stretchr/testify/assert"
)

type testLayout struct {
	label string
}

func (t testLayout) Layout() fyne.CanvasObject {
	return widget.NewLabel(t.label)
}

func TestScreenRouter(t *testing.T) {
	t.Parallel()
	test.NewApp()

	main := testLayout{"main"}
	router := logic.NewScreenRouter(nil)
	router.Reset(main)

	assert.Equal(t, main.label, router.Container().Objects[0].(*widget.Label).Text)

	secondary := testLayout{"secondary"}
	router.Push(secondary)
	assert.Equal(t, secondary.label, router.Container().Objects[0].(*widget.Label).Text)

	router.Pop()
	assert.Equal(t, main.label, router.Container().Objects[0].(*widget.Label).Text)
	router.Pop()
	assert.Equal(t, main.label, router.Container().Objects[0].(*widget.Label).Text)
}
