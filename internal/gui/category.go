package gui

import (
	"errors"
	"log"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const categoryMinLength = 4

type category struct {
	categoriesStore categoriesStore

	router router
	canvas fyne.Canvas
}

func NewCategory(categoriesStore categoriesStore, router router, canvas fyne.Canvas) *category {
	return &category{
		categoriesStore: categoriesStore,

		router: router,
		canvas: canvas,
	}
}

func (c category) Layout() fyne.CanvasObject {
	return container.NewBorder(c.top(), c.bottom(), nil, nil, c.center())
}

func (c category) top() *fyne.Container {
	input := widget.NewEntry()

	addCategory := func() {
		name := input.Text
		err := c.categoriesStore.Add(name)
		if err != nil {
			log.Printf("could not add new category: %v", err)
			return
		}
		c.router.Reload()
	}

	input.OnSubmitted = func(s string) {
		addCategory()
	}

	right := widget.NewButtonWithIcon("", theme.ContentAddIcon(), func() {
		addCategory()
	})

	return container.NewBorder(nil, nil, nil, right, input)
}

func (c category) center() *fyne.Container {
	categories := c.categoriesStore.GetAll()

	sort.Slice(categories, func(i, j int) bool {
		return categories[i].Name < categories[j].Name
	})

	rows := container.NewVBox()
	for _, cat := range categories {
		cat := cat
		right := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() {
			body := widget.NewLabel("This will remove category \nand all its decks for eternity.")
			body.Alignment = fyne.TextAlignCenter
			content := container.NewVBox(
				container.NewCenter(widget.NewLabel("Are you sure?")),
				container.NewCenter(body),
				container.NewGridWithRows(1,
					widget.NewButton("No", func() {
						c.router.CleanPopup()
					}),
					widget.NewButton("Yes", func() {
						err := c.categoriesStore.Remove(cat.ID)
						if err != nil {
							log.Printf("could not delete category: %v", err)
						}
						c.router.Reload()
						c.router.CleanPopup()
					}),
				),
			)
			c.router.SetPopup(content)
		})

		categoryName := container.NewMax()
		entry := widget.NewEntry()
		btn := widget.NewButton(cat.Name, nil)
		btn.Importance = widget.LowImportance
		btn.OnTapped = func() {
			entry.SetText(cat.Name)
			categoryName.RemoveAll()
			categoryName.Add(entry)
			c.canvas.Focus(entry)
			entry.CursorColumn = len(cat.Name)
		}

		entry.OnSubmitted = func(s string) {
			if err := c.categoriesStore.Rename(cat.ID, s); err != nil {
				log.Printf("could not rename category: %v", err)
				return
			}

			btn.SetText(s)
			categoryName.RemoveAll()
			categoryName.Add(btn)
		}
		categoryName.Add(btn)

		row := container.NewBorder(nil, nil, nil, right, categoryName)
		rows.Add(row)
	}
	return container.NewBorder(nil, nil, nil, nil, rows)
}

func (c category) bottom() *fyne.Container {
	returnBtn := widget.NewButtonWithIcon("", theme.NavigateBackIcon(), func() {
		c.router.Pop()
		c.router.Reload()
	})
	return container.NewBorder(nil, nil, returnBtn, nil, container.NewMax())
}

func categoryNameValidator(s string) error {
	if len(s) < categoryMinLength {
		return errors.New("too short")
	}
	return nil
}
