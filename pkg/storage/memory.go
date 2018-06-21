package storage

import (
	"errors"
)

// MemoryStorage is an in-memory no-op implementation of the sup.Storer interface.
type MemoryStorage struct {
	// Storing the data in a basic map.
	data map[string]string
}

// NewInMemoryStorage initializes the storage.
func NewInMemoryStorage() *MemoryStorage {
	return &MemoryStorage{data: make(map[string]string)}
}

// Load is a no-op.
func (m *MemoryStorage) Load() error { return nil }

// Save is a no-op.
func (m *MemoryStorage) Save() error { return nil }

// Put adds a key-value pair.
func (m *MemoryStorage) Put(key, value string) {
	m.data[key] = value
}

// Get returns the value or an error if not.
func (m *MemoryStorage) Get(key string) (string, error) {
	if value, ok := m.data[key]; ok {
		return value, nil
	}
	return "", errors.New("not found")
}
