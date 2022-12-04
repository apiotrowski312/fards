package store_test

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/apiotrowski312/fards/internal/logic/store"
	"github.com/stretchr/testify/assert"
)

func TestCategories(t *testing.T) {
	cat := store.NewCategory(test.NewApp().Storage())

	assert.Equal(t, len(cat.GetAll()), 0)
	err := cat.Add("cat1")
	assert.Equal(t, err, nil)
	assert.Equal(t, len(cat.GetAll()), 1)

	err = cat.Remove("not-real-id")
	assert.Equal(t, err, nil)

	assert.Equal(t, len(cat.GetAll()), 1)
	err = cat.Remove(cat.GetAll()[0].ID)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(cat.GetAll()), 0)
}
