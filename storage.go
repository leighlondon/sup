package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Load the data stored in a file.
func load(filename string) map[string]string {
	// Load the data in as bytes, and then into a data structure.
	bytes := readFile(filename)
	data := loadItems(bytes)
	return data
}

// Save the data to a file.
func save(filename string, data map[string]string) {
	// Marshal the data structure into bytes, and then save to a file.
	bytes := makeBytes(data)
	saveFile(filename, bytes)
}

// Read a file by filename into memory as bytes.
func readFile(filename string) []byte {

	// Read in the file as bytes using ioutil.
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		printWarn("unable to read file")
		os.Exit(1)
	}

	return bytes
}

// Save some bytes to a file.
func saveFile(filename string, data []byte) {

	// Need to provide the permissions even if file already exists.
	// The default permissions to be used will be 0644.
	err := ioutil.WriteFile(filename, data, 0644)

	if err != nil {
		printWarn("problem saving file")
		os.Exit(1)
	}
}

// Load bytes into a data structure.
func loadItems(bytes []byte) map[string]string {

	// The data structure internally is a map from strings to strings.
	data := make(map[string]string)
	// Unmarshal the data from JSON-as-bytes into the map.
	err := json.Unmarshal(bytes, &data)

	if err != nil {
		printWarn("unable to load data")
		os.Exit(1)
	}

	return data
}

// Convert a data structure into byte format.
func makeBytes(data map[string]string) []byte {

	// Marshal the map into JSON, as bytes.
	bytes, err := json.Marshal(data)

	if err != nil {
		printWarn("unable to marshal data")
		os.Exit(1)
	}

	return bytes
}
