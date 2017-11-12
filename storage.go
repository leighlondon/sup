package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JSONStorage struct {
	data     map[string]string
	filename string
}

// Load loads the data stored in a file.
func (s *JSONStorage) Load() error {
	// Load the data in as bytes, and then into a data structure.
	bytes, err := readFile(s.filename)
	if err != nil {
		return err
	}
	s.data = loadItems(bytes)
	return nil
}

// Save saves the data to a file.
func (s *JSONStorage) Save() error {
	// Marshal the data structure into bytes, and then save to a file.
	bytes, err := makeBytes(s.data)
	if err != nil {
		return err
	}
	err = saveFile(s.filename, bytes)
	if err != nil {
		return err
	}
	return nil
}

// Read a file by filename into memory as bytes.
func readFile(filename string) ([]byte, error) {
	// Read in the file as bytes using ioutil.
	return ioutil.ReadFile(filename)
}

// Save some bytes to a file.
func saveFile(filename string, data []byte) error {
	// Need to provide the permissions even if file already exists.
	// The default permissions to be used will be 0644.
	return ioutil.WriteFile(filename, data, 0644)
}

// Load bytes into a data structure.
func loadItems(bytes []byte) map[string]string {
	// The data structure internally is a map from strings to strings.
	data := make(map[string]string)
	// Fail early if there's no data provided; just return the empty map.
	if bytes == nil {
		return data
	}
	// Unmarshal the data from JSON-as-bytes into the map.
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to load data")
		os.Exit(1)
	}
	return data
}

// Convert a data structure into byte format.
func makeBytes(data map[string]string) ([]byte, error) {
	// Marshal the map into JSON, as bytes.
	return json.Marshal(data)
}
