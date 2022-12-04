package store_test

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/apiotrowski312/fards/internal/logic/store"
	"github.com/apiotrowski312/fards/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestDecksListStore(t *testing.T) {
	t.Parallel()

	deckListStore := store.NewDecks(test.NewApp().Storage())

	list := deckListStore.Get()
	assert.Equal(t, len(list), 0)

	entry, err := deckListStore.Upsert(models.DecksListEntry{Name: "test", CategoryID: "test"})
	assert.Equal(t, err, nil)
	assert.Equal(t, entry.Name, "test")
	assert.Equal(t, entry.CategoryID, "test")

	list = deckListStore.Get()
	assert.Equal(t, len(list), 1)

	newEntry, err := deckListStore.Upsert(models.DecksListEntry{ID: entry.ID, Name: "next_test", CategoryID: "next_test"})
	assert.Equal(t, err, nil)
	assert.Equal(t, newEntry.FileName, entry.FileName)
	assert.Equal(t, newEntry.ID, newEntry.ID)
	assert.Equal(t, newEntry.Name, "next_test")
	assert.Equal(t, newEntry.CategoryID, "next_test")

	err = deckListStore.Delete(newEntry.ID)
	assert.Equal(t, err, nil)

	list = deckListStore.Get()
	assert.Equal(t, len(list), 0)
}
