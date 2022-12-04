package store

import (
	"encoding/json"
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/apiotrowski312/fards/internal/models"
)

const categoriesFile = "categories.json"

var (
	ErrCategoryDuplicated = errors.New("category duplicated")
	ErrCategoryEmpty      = errors.New("category empty")
)

type categoryStore struct {
	storage fyne.Storage

	categories models.Categories
}

func NewCategory(storage fyne.Storage) *categoryStore {
	categoryStore := &categoryStore{
		storage: storage,
	}

	// FIXME: Do not ignore error
	categoryStore.load()

	return categoryStore
}

func (s categoryStore) GetAll() models.Categories {
	return s.categories
}

func (s categoryStore) Rename(id, newName string) error {
	for _, cat := range s.categories {
		if cat.ID == id {
			cat.Name = newName
			break
		}
	}
	return s.save()
}

func (s *categoryStore) Add(name string) error {
	if name == "" {
		return ErrCategoryEmpty
	}

	for _, c := range s.categories {
		if c.Name == name {
			return ErrCategoryDuplicated
		}
	}

	s.categories = append(s.categories, models.NewCategory(name))
	return s.save()
}

func (s *categoryStore) Remove(id string) error {
	index := -1
	for i, c := range s.categories {
		if c.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return nil
	}

	dl := NewDecks(s.storage)
	dl.deleteByCategoryID(id)
	s.categories = append(s.categories[:index], s.categories[index+1:]...)
	return s.save()
}

func (s *categoryStore) load() error {
	bytes, err := load(s.storage, categoriesFile)
	if err != nil {
		return fmt.Errorf("could not load categories: %w", err)
	}

	categories := make(models.Categories, 0)
	if err := json.Unmarshal(bytes, &categories); err != nil {
		return fmt.Errorf("could not unmarshal categories: %w", err)
	}

	s.categories = categories
	return nil
}

func (s categoryStore) save() error {
	return save(s.storage, categoriesFile, s.categories)
}
