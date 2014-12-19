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

func readFile(filename string) []byte {

	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		printWarn("unable to read file")
		os.Exit(1)
	}

	return bytes
}

func saveFile(filename string, data []byte) {

	err := ioutil.WriteFile(filename, data, 0644)

	if err != nil {
		printWarn("problem saving file")
		os.Exit(1)
	}
}

func loadItems(bytes []byte) map[string]string {

	data := make(map[string]string)

	err := json.Unmarshal(bytes, &data)

	if err != nil {
		printWarn("unable to load data")
		os.Exit(1)
	}

	return data
}

func makeBytes(data map[string]string) []byte {

	bytes, err := json.Marshal(data)

	if err != nil {
		printWarn("unable to marshal data")
		os.Exit(1)
	}

	return bytes
}
