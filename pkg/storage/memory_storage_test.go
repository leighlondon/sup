package storage

import (
	"testing"
)

func inMemoryDouble() *MemoryStorage {
	return NewInMemoryStorage("test")
}

func TestMemoryStorageBasicOps(t *testing.T) {
	m := inMemoryDouble()
	m.Put("hello", "world")
	value, err := m.Get("hello")
	if err != nil {
		t.Error("should not error")
	}
	if value != "world" {
		t.Errorf("expected 'world', got %s", value)
	}
}

func TestMemoryGetErrors(t *testing.T) {
	m := inMemoryDouble()
	if _, err := m.Get("absurd"); err == nil {
		t.Error("should have returned error")
	}
}

func TestMemoryLoad(t *testing.T) {
	m := inMemoryDouble()
	if m.Load() != nil {
		t.Error("load should not return error")
	}
}

func TestMemorySave(t *testing.T) {
	m := inMemoryDouble()
	if m.Save() != nil {
		t.Error("save should not return error")
	}
}

func TestMemoryFilename(t *testing.T) {
	m := inMemoryDouble()
	if m.Filename() == "" {
		t.Error("should be non-empty string")
	}
}

func TestMemoryAll(t *testing.T) {
	m := inMemoryDouble()
	m.Put("hello", "world")
	if len(m.All()) != 1 {
		t.Error("all should return data")
	}
}
