package gui

import (
	"fyne.io/fyne/v2"
	"github.com/apiotrowski312/fards/internal/logic"
	"github.com/apiotrowski312/fards/internal/models"
)

//go:generate mockgen -source gui.go -destination mocks_test.go -package gui_test
type decksStore interface {
	Get() models.DecksList
	Delete(id string) error
	Upsert(deck models.DecksListEntry) (models.DecksListEntry, error)
}

type categoriesStore interface {
	GetAll() models.Categories
	Remove(id string) error
	Add(name string) error
	Rename(id, newName string) error
}

type router interface {
	CleanPopup()
	SetPopup(content fyne.CanvasObject)
	Pop()
	Push(screen logic.Layoutable)
	Reload()
}

type cardsStore interface {
	GetCard() (string, string)
	Next() bool
	Previous() bool
	Shuffle()
	Update(front, back string)
	Add(front, back string)
	RemoveCurrentCard()
	GetStats() (int, int)
}
