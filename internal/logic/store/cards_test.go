package store_test

import (
	"fmt"
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/apiotrowski312/fards/internal/logic/store"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestDeckStore_CRUD(t *testing.T) {
	t.Parallel()

	deckStore := store.NewCards(test.NewApp().Storage(), uuid.NewString()+".csv")

	q, a := deckStore.GetCard()
	assert.Equal(t, "", q)
	assert.Equal(t, "", a)

	deckStore.Add("question", "answer")
	q, a = deckStore.GetCard()
	assert.Equal(t, "", q)
	assert.Equal(t, "", a)

	deckStore.Next()
	q, a = deckStore.GetCard()
	assert.Equal(t, "question", q)
	assert.Equal(t, "answer", a)

	deckStore.Update("q", "a")
	q, a = deckStore.GetCard()
	assert.Equal(t, "q", q)
	assert.Equal(t, "a", a)

	deckStore.RemoveCurrentCard()
	q, a = deckStore.GetCard()
	assert.Equal(t, "", q)
	assert.Equal(t, "", a)
}

func TestDeckStore_NextPrev(t *testing.T) {
	t.Parallel()

	deckStore := store.NewCards(test.NewApp().Storage(), uuid.NewString()+".csv")
	deckStore.Update("1", "1")
	deckStore.Add("2", "2")
	deckStore.Add("3", "3")

	q, a := deckStore.GetCard()
	assert.Equal(t, "1", q)
	assert.Equal(t, "1", a)

	deckStore.Previous()
	q, a = deckStore.GetCard()
	assert.Equal(t, "1", q)
	assert.Equal(t, "1", a)

	deckStore.Next()
	deckStore.Next()
	q, a = deckStore.GetCard()
	assert.Equal(t, "3", q)
	assert.Equal(t, "3", a)

	deckStore.Next()
	q, a = deckStore.GetCard()
	assert.Equal(t, "3", q)
	assert.Equal(t, "3", a)

	deckStore.Previous()
	q, a = deckStore.GetCard()
	assert.Equal(t, "2", q)
	assert.Equal(t, "2", a)
}

func TestDeckStore_OpenFile(t *testing.T) {
	t.Parallel()

	file := uuid.NewString() + ".csv"

	deckStore := store.NewCards(test.NewApp().Storage(), file)
	deckStore.Add("question", "answer")
	_ = store.NewCards(test.NewApp().Storage(), file)
}

func TestDeckStore_Shuffle(t *testing.T) {
	t.Parallel()

	deckStore := store.NewCards(test.NewApp().Storage(), uuid.NewString()+".csv")
	for i := 0; i <= 5; i++ {
		deckStore.Add(fmt.Sprint(i), "answer")
	}

	deckStore.Shuffle()
	wasShuffled := false
	for i := 0; i <= 5; i++ {
		deckStore.Next()
		q, _ := deckStore.GetCard()

		if q != fmt.Sprint(i) {
			wasShuffled = true
			break
		}
	}
	assert.Equal(t, wasShuffled, true)
}
