package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"github.com/apiotrowski312/fards/internal/models"
	"github.com/google/uuid"
)

const decksListFile = "deck_list.json"

type decksStore struct {
	storage fyne.Storage

	list models.DecksList
}

func NewDecks(storage fyne.Storage) *decksStore {
	listStore := &decksStore{
		storage: storage,
	}

	// FIXME: Do not ignore error
	listStore.load()

	return listStore
}

func (s decksStore) Get() models.DecksList {
	return s.list
}

func (s *decksStore) Upsert(deck models.DecksListEntry) (models.DecksListEntry, error) {
	if deck.CategoryID == "" {
		return models.DecksListEntry{}, errors.New("category id cannot be empty")
	}

	index := -1
	for i, r := range s.list {
		if r.ID == deck.ID {
			index = i

			// Overwrite filename to be sure it is not messed up
			deck.FileName = r.FileName
			break
		}
	}

	if index == -1 {
		deck.ID = uuid.NewString()
		deck.FileName = deck.ID + ".csv"
		s.list = append(s.list, deck)
	} else {
		s.list[index] = deck
	}

	if err := s.save(); err != nil {
		return models.DecksListEntry{}, fmt.Errorf("could not save entry: %w", err)
	}
	return deck, nil
}

func (s *decksStore) Delete(id string) error {
	index := -1
	for i, r := range s.list {
		if r.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return nil
	}

	// TODO: We are not removing list from device

	s.list = append(s.list[:index], s.list[index+1:]...)
	return s.save()
}

func (s *decksStore) load() error {
	bytes, err := load(s.storage, decksListFile)
	if err != nil {
		return fmt.Errorf("could not load deck: %w", err)
	}

	var decks models.DecksList
	if err := json.Unmarshal(bytes, &decks); err != nil {
		return fmt.Errorf("could not unmarshal deck: %w", err)
	}

	s.list = decks
	return nil
}

func (s decksStore) save() error {
	return save(s.storage, decksListFile, s.list)
}

func (s *decksStore) deleteByCategoryID(categoryID string) {
	ids := make([]string, 0)
	for _, r := range s.list {
		if r.CategoryID == categoryID {
			ids = append(ids, r.ID)
		}
	}
	for _, id := range ids {
		if err := s.Delete(id); err != nil {
			log.Printf("could not delete by category: %v", err)
		}
	}
}
