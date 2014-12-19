package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func load(filename string) map[string]string {

	bytes := readFile(filename)
	data := loadItems(bytes)

	return data
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
