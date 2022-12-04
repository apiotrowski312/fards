package models

import "github.com/google/uuid"

type Category struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewCategory(name string) *Category {
	return &Category{
		Name: name,
		ID:   uuid.NewString(),
	}
}

type Categories []*Category
