package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/apiotrowski312/fards/internal/gui"
	"github.com/apiotrowski312/fards/internal/logic"
	"github.com/apiotrowski312/fards/internal/logic/store"
)

func main() {
	a := app.NewWithID("fards")
	w := a.NewWindow("fards")
	c := w.Canvas()
	router := logic.NewScreenRouter(c)
	home := gui.NewHome(store.NewDecks(a.Storage()), store.NewCategory(a.Storage()), router, c, a.Storage())

	router.Reset(home)

	splash := gui.MakeSplash()
	cntr := container.NewMax(
		router.Container(),
		splash,
	)
	w.SetContent(
		cntr,
	)

	go gui.FadeSplash(splash)

	w.SetPadded(true)

	w.ShowAndRun()
}
