package storage

import (
	"encoding/json"
	"io/ioutil"
)

// FileBackedStorage implements the sup.Storer interface.
type FileBackedStorage struct {
	data     map[string]string
	filename string
}

// New initializes a new instance of FileBackedStorage.
func New(filename string) (*FileBackedStorage, error) {
	data, err := loadFile(filename)
	return &FileBackedStorage{data: data, filename: filename}, err
}

// Filename returns the filename of the storage.
func (s *FileBackedStorage) Filename() string {
	return s.filename
}

// Load updates the state from the backing store.
func (s *FileBackedStorage) Load() error {
	data, err := loadFile(s.filename)
	if err != nil {
		return err
	}
	s.data = data
	return nil
}

// Save the state to the backing store.
func (s *FileBackedStorage) Save() error {
	return saveFile(s.filename, s.data)
}

func loadFile(filename string) (map[string]string, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	data := make(map[string]string)
	err = json.Unmarshal(b, &data)
	return data, err
}

func saveFile(filename string, data map[string]string) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, b, 0644)
}
