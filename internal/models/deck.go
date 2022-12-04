package models

type DecksListEntry struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	CategoryID string `json:"category,omitempty"`
	FileName   string `json:"file_name,omitempty"`
}

type DecksList []DecksListEntry
