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

func TestCategory(t *testing.T) {
	t.Parallel()

	testCase := []struct {
		name string

		prepare func(c fyne.Canvas)
	}{
		{
			name:    "render",
			prepare: func(c fyne.Canvas) {},
		},
		{
			name: "remove",
			prepare: func(c fyne.Canvas) {
				test.TapCanvas(c, fyne.NewPos(280, 60))
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
			categoriesStore := NewMockcategoriesStore(mockCtrl)

			categoriesStore.EXPECT().GetAll().Return(models.Categories{
				{ID: "1", Name: "Polish"},
				{ID: "2", Name: "German"},
			})

			w := a.NewWindow("category")
			c := w.Canvas()

			router := logic.NewScreenRouter(w.Canvas())
			category := gui.NewCategory(categoriesStore, router, c)
			router.Reset(category)

			w.SetContent(router.Container())
			w.Resize(fyne.NewSize(300, 400))

			tc.prepare(c)

			test.AssertImageMatches(t, "category_"+strings.ReplaceAll(tc.name, " ", "_")+".png", c.Capture())
		})
	}
}
