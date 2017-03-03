package main

import (
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
	marshalled, _ := makeBytes(data)
	unmarshalled := loadItems(marshalled)
	// Make sure that it's equal.
	if reflect.DeepEqual(unmarshalled, data) == false {
		t.Error("map not decoded correctly")
	}
}
