package gui_test

import (
	"strings"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"github.com/apiotrowski312/fards/internal/gui"
	"github.com/apiotrowski312/fards/internal/logic"
	"github.com/apiotrowski312/fards/internal/models"
	"github.com/golang/mock/gomock"
)

func TestHome(t *testing.T) {
	t.Parallel()

	testCase := []struct {
		name string

		prepare func(c fyne.Canvas)
	}{
		{
			name: "unfolded home",
			prepare: func(c fyne.Canvas) {
				test.TapCanvas(c, fyne.NewPos(100, 20))
			},
		},
		{
			name: "create popup",
			prepare: func(c fyne.Canvas) {
				test.TapCanvas(c, fyne.NewPos(280, 20))
			},
		},
		{
			name: "rename popup",
			prepare: func(c fyne.Canvas) {
				test.TapCanvas(c, fyne.NewPos(100, 20))
				test.TapCanvas(c, fyne.NewPos(280, 110))
				test.TapCanvas(c, fyne.NewPos(280, 175))
			},
		},
		{
			name: "delete popup",
			prepare: func(c fyne.Canvas) {
				test.TapCanvas(c, fyne.NewPos(100, 20))
				test.TapCanvas(c, fyne.NewPos(280, 110))
				test.TapCanvas(c, fyne.NewPos(280, 185))
			},
		},
	}
	for _, tc := range testCase {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			a := test.NewApp()
			defer a.Quit()

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			decksStore := NewMockdecksStore(mockCtrl)
			categoriesStore := NewMockcategoriesStore(mockCtrl)

			decksStore.EXPECT().Get().Return(models.DecksList{
				{ID: "1", Name: "Salads", CategoryID: "1"},
				{ID: "2", Name: "Soups", CategoryID: "1"},
				{ID: "3", Name: "Salads", CategoryID: "2"},
				{ID: "4", Name: "Soups", CategoryID: "2"},
			})
			categoriesStore.EXPECT().GetAll().Return(models.Categories{
				{ID: "1", Name: "Polish"},
				{ID: "2", Name: "German"},
			})

			w := a.NewWindow("home")
			c := w.Canvas()

			router := logic.NewScreenRouter(w.Canvas())
			home := gui.NewHome(decksStore, categoriesStore, router, w.Canvas(), a.Storage())
			router.Reset(home)

			w.SetContent(router.Container())
			w.Resize(fyne.NewSize(300, 400))

			tc.prepare(c)

			test.AssertImageMatches(t, "home_"+strings.ReplaceAll(tc.name, " ", "_")+".png", c.Capture())
		})
	}
}
