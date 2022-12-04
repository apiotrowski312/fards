package gui

import (
	"log"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/apiotrowski312/fards/internal/logic/store"
	"github.com/apiotrowski312/fards/internal/models"
)

type home struct {
	decksStore      decksStore
	categoriesStore categoriesStore

	router  router
	canvas  fyne.Canvas
	storage fyne.Storage
}

func NewHome(decksStore decksStore, categoriesStore categoriesStore, router router, canvas fyne.Canvas, storage fyne.Storage) *home {
	return &home{
		decksStore:      decksStore,
		categoriesStore: categoriesStore,

		router:  router,
		canvas:  canvas,
		storage: storage,
	}
}

func (h home) Layout() fyne.CanvasObject {
	list := h.list()
	bottom := h.bottom()

	screen := container.NewBorder(nil, bottom, nil, nil, list)
	return container.NewMax(screen)
}

func (h home) list() fyne.CanvasObject {
	decks := h.decksStore.Get()
	categories := h.categoriesStore.GetAll()

	listOfDecks := container.NewVBox()
	categoryContainer := make(map[string][2]*fyne.Container)
	for _, category := range categories {
		deckRowsContainer := container.NewVBox()
		deckRowsContainer.Hide()

		label := h.categoryLabelRow(*category, deckRowsContainer, listOfDecks)
		categoryContainer[category.ID] = [2]*fyne.Container{
			label, deckRowsContainer,
		}
	}

	for _, deck := range decks {
		categoryList, isOk := categoryContainer[deck.CategoryID]
		if !isOk {
			continue
		}
		categoryList[1].Add(h.deckRow(deck))
	}

	sort.Slice(categories, func(i, j int) bool {
		return categories[i].Name < categories[j].Name
	})

	for _, category := range categories {
		row, isOk := categoryContainer[category.ID]
		if !isOk {
			continue
		}
		listOfDecks.Add(row[0])
		listOfDecks.Add(row[1])
	}

	return listOfDecks
}

func (h home) categoryLabelRow(category models.Category, deckRowsContainer, parentContainer *fyne.Container) *fyne.Container {
	foldBtn := widget.NewButtonWithIcon("", theme.MenuDropDownIcon(), nil)
	foldBtn.OnTapped = func() {
		if deckRowsContainer.Hidden {
			foldBtn.SetIcon(theme.MenuDropUpIcon())
			deckRowsContainer.Show()
			parentContainer.Refresh()
			return
		}
		foldBtn.SetIcon(theme.MenuDropDownIcon())
		deckRowsContainer.Hide()
		parentContainer.Refresh()
	}

	addBtn := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		deckNameEntry := widget.NewEntry()
		deckNameEntry.SetPlaceHolder("Deck name")

		next := func() {
			dl, err := h.decksStore.Upsert(models.DecksListEntry{
				Name:       deckNameEntry.Text,
				CategoryID: category.ID,
			})

			if err != nil {
				log.Printf("could not upsert deck: %v", err)
			}

			h.router.Reload()
			h.router.Push(NewEditor(store.NewCards(h.storage, dl.FileName), h.router, h.canvas))
			h.router.CleanPopup()
		}

		content := container.NewVBox(
			deckNameEntry,
			container.NewGridWithRows(1,
				widget.NewButton("Cancel", func() {
					h.router.CleanPopup()
				}),
				widget.NewButton("Next", func() {
					next()
				}),
			),
		)

		deckNameEntry.OnSubmitted = func(s string) { next() }

		h.router.SetPopup(content)
	})

	label := container.NewBorder(nil, nil, nil,
		addBtn,
		container.NewHBox(widget.NewLabel(category.Name), foldBtn),
	)
	return label
}

func (h home) deckRow(deck models.DecksListEntry) *fyne.Container {
	labelBtn := widget.NewButton(deck.Name, nil)
	dropdown := widget.NewSelect([]string{"Edit cards", "Rename", "Delete"}, nil)
	center := container.NewBorder(nil, nil, nil, dropdown, labelBtn)

	dropdown.Alignment = fyne.TextAlignLeading
	dropdown.PlaceHolder = "Action"
	dropdown.OnChanged = func(option string) {
		switch option {
		case "Edit cards":
			h.router.Push(NewEditor(store.NewCards(h.storage, deck.FileName), h.router, h.canvas))
			return
		case "Rename":
			deckNameEntry := widget.NewEntry()
			deckNameEntry.SetPlaceHolder("Deck name")
			deckNameEntry.SetText(deck.Name)

			update := func() {
				deck.Name = deckNameEntry.Text
				_, err := h.decksStore.Upsert(deck)
				if err != nil {
					// TODO: Add popup with info
					log.Printf("Could not upsert deck: %v", err)
					return
				}

				h.router.CleanPopup()

				labelBtn.SetText(deck.Name)
			}

			content := container.NewVBox(
				deckNameEntry,
				container.NewGridWithRows(1,
					widget.NewButton("Cancel", func() {
						h.router.CleanPopup()
					}),
					widget.NewButton("Update", func() {
						update()
					}),
				),
			)

			deckNameEntry.OnSubmitted = func(s string) { update() }

			h.router.SetPopup(content)
		case "Delete":
			content := container.NewVBox(
				container.NewCenter(widget.NewLabel("Are you sure?")),
				container.NewCenter(widget.NewLabel("This will remove deck for eternity.")),
				container.NewGridWithRows(1,
					widget.NewButton("No", func() {
						h.router.CleanPopup()
					}),
					widget.NewButton("Yes", func() {
						err := h.decksStore.Delete(deck.ID)
						if err != nil {
							log.Printf("could not delete deck: %v", err)
						}

						center.Hide()
						h.router.CleanPopup()
					}),
				),
			)
			h.router.SetPopup(content)
		}
	}

	labelBtn.OnTapped = func() {
		cards := store.NewCards(h.storage, deck.FileName)
		h.router.Push(NewViewer(cards, h.router, h.canvas))
	}

	labelBtn.Alignment = widget.ButtonAlignLeading
	labelBtn.Importance = widget.LowImportance

	return center
}

func (h home) bottom() *fyne.Container {
	categoryBtn := widget.NewButton("Categories", func() {
		category := NewCategory(h.categoriesStore, h.router, h.canvas)
		h.router.Push(category)
	})

	return container.NewGridWithRows(1, layout.NewSpacer(), categoryBtn, layout.NewSpacer())
}
