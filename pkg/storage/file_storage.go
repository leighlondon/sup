package storage

import (
	"encoding/json"
	"errors"
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

// All returns all key-value pairs.
func (s *FileBackedStorage) All() map[string]string {
	return s.data
}

// Filename returns the filename of the storage.
func (s *FileBackedStorage) Filename() string {
	return s.filename
}

// Put adds a key-value pair.
func (s *FileBackedStorage) Put(key, value string) {
	s.data[key] = value
}

// Get returns the value or an error if not.
func (s *FileBackedStorage) Get(key string) (string, error) {
	if value, ok := s.data[key]; ok {
		return value, nil
	}
	return "", errors.New("not found")
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
