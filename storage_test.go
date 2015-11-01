package main

import (
	"os"
	"reflect"
	"testing"
)

func TestMapToBytesToMapProducesSameMap(t *testing.T) {
	// If we take an input map and convert to bytes and then
	// convert back to the map, it should match to be exactly the same.
	data := make(map[string]string)
	data["hello"] = "world"
	data["message"] = "response"
	// Round loop test.
	marshalled := makeBytes(data)
	unmarshalled := loadItems(marshalled)
	// Make sure that it's equal.
	if reflect.DeepEqual(unmarshalled, data) == false {
		t.Error("map not decoded correctly")
	}
}

func TestSaveToLoadProducesSameData(t *testing.T) {
	// Saving data and then loading it should produce the same data.
	data := make(map[string]string)
	data["hello2"] = "world2"
	data["dolla"] = "dolla bills"
	filename := "sup_test_filename"
	// Save then load it.
	SaveData(filename, data)
	loaded := LoadData(filename)
	// Make sure that it's equal.
	if reflect.DeepEqual(data, loaded) == false {
		t.Error("saved data different from loaded")
	}
	// Clear the temporary file.
	os.Remove(filename)
}
