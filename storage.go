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
	bytes := readFile(s.filename)
	data := loadItems(bytes)
	s.data = data
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

// Check if a file does not exist.
func fileNotExists(filename string) bool {
	// Just stat the filename and get the error.
	_, err := os.Stat(filename)
	if err != nil {
		// Check if the error says that it doesn't exist.
		if os.IsNotExist(err) {
			return true
		}
	}
	// If there was no errors the file is considered to exist.
	return false
}

// Read a file by filename into memory as bytes.
func readFile(filename string) []byte {
	// If the file doesn't exist there's no data to read.
	if fileNotExists(filename) == true {
		return nil
	}
	// Read in the file as bytes using ioutil.
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to read file")
		os.Exit(1)
	}
	return bytes
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
