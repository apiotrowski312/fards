package gui_test

import (
	"strings"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/apiotrowski312/fards/internal/gui"
	"github.com/apiotrowski312/fards/internal/logic"
	"github.com/golang/mock/gomock"
)

func TestEditor(t *testing.T) {
	t.Parallel()

	testCase := []struct {
		name string

		prepare func(c fyne.Canvas)
	}{
		{
			name:    "render",
			prepare: func(c fyne.Canvas) {},
		},
	}
	for _, tc := range testCase {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			a := test.NewApp()
			defer a.Quit()

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			cardsStore := NewMockcardsStore(mockCtrl)

			cardsStore.EXPECT().GetCard().Return("front", "back")
			cardsStore.EXPECT().Next().Return(false)
			cardsStore.EXPECT().GetStats().Return(1, 11).Times(2)

			w := a.NewWindow("editor")
			c := w.Canvas()

			router := logic.NewScreenRouter(w.Canvas())
			editor := gui.NewEditor(cardsStore, router, c)
			router.Reset(editor)

			w.SetContent(router.Container())
			w.Resize(fyne.NewSize(300, 400))

			tc.prepare(c)

			test.AssertImageMatches(t, "editor_"+strings.ReplaceAll(tc.name, " ", "_")+".png", c.Capture())
		})
	}
}
