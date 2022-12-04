package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
)

func createWriter(s fyne.Storage, fileName string) (io.WriteCloser, error) {
	writer, err := s.Save(fileName)
	if err != nil {
		if errors.Is(err, storage.ErrNotExists) {
			writer, err = s.Create(fileName)
		}

		if err != nil {
			return nil, fmt.Errorf("could not save/create file: %w", err)
		}
	}
	return writer, err
}

func load(storage fyne.Storage, file string) ([]byte, error) {
	reader, err := storage.Open(file)
	if err != nil {
		files := storage.List()
		for _, fileName := range files {
			// If file exist it means there were other problem
			if fileName == file {
				return nil, fmt.Errorf("could not open file: %w", err)
			}
		}
		return []byte{}, nil
	}
	defer reader.Close()

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}
	return bytes, nil
}

func save(storage fyne.Storage, file string, toSave interface{}) error {
	writer, err := createWriter(storage, file)
	if err != nil {
		return fmt.Errorf("could not create writer: %w", err)
	}
	defer writer.Close()

	bytes, err := json.Marshal(toSave)
	if err != nil {
		return fmt.Errorf("could not marshal interface: %w", err)
	}

	if _, err := writer.Write(bytes); err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	return nil
}
