package store

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
)

type cardsStore struct {
	storage fyne.Storage

	index int
	cards [][2]string

	fileName string
}

// NewCards will load a deck of cards. If deck does not exists it will load with one empty card.
func NewCards(storage fyne.Storage, fileName string) *cardsStore {
	cardsStore := &cardsStore{
		storage:  storage,
		fileName: fileName,

		index: 0,
	}

	cardsStore.load()

	if len(cardsStore.cards) == 0 {
		cardsStore.cards = [][2]string{{}}
	}

	return cardsStore
}

func (s *cardsStore) GetStats() (int, int) {
	return s.index + 1, len(s.cards)
}

func (s *cardsStore) GetCard() (string, string) {
	if s.index < 0 || s.index >= len(s.cards) {
		return "", ""
	}

	card := s.cards[s.index]
	return card[0], card[1]
}

func (s *cardsStore) Next() bool {
	if s.index >= len(s.cards)-1 {
		return false
	}
	s.index++
	return true
}

func (s *cardsStore) Previous() bool {
	if s.index <= 0 {
		return false
	}
	s.index--
	return true
}

func (s *cardsStore) Shuffle() {
	cards := s.cards
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })

	s.cards = cards
}

func (s *cardsStore) Update(front, back string) {
	s.cards[s.index] = [2]string{front, back}

	s.save()
}

func (s *cardsStore) Add(front, back string) {
	s.cards = append(s.cards, [2]string{front, back})

	s.save()
}

func (s *cardsStore) RemoveCurrentCard() {
	if s.index == 0 && len(s.cards) > 0 {
		s.cards = [][2]string{{}}
		s.save()
		return
	}
	s.cards = append(s.cards[:s.index], s.cards[s.index+1:]...)
	s.save()
	s.index--
}

func (s *cardsStore) load() error {
	bytes, err := load(s.storage, s.fileName)
	if err != nil {
		return fmt.Errorf("could not load cards: %w", err)
	}

	cards := make([][2]string, 0)
	if err := json.Unmarshal(bytes, &cards); err != nil {
		return fmt.Errorf("could not unmarshal cards: %w", err)
	}

	s.cards = cards
	return nil
}

func (s cardsStore) save() error {
	return save(s.storage, s.fileName, s.cards)
}
