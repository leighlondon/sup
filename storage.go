package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func load(filename string) map[string]string {

	data := readFile(filename)
	items := loadItems(data)

	return items
}

func readFile(filename string) []byte {

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		printWarn("unable to read file")
		os.Exit(1)
	}

	return data
}

func loadItems(input []byte) map[string]string {

	data := make(map[string]string)

	err := json.Unmarshal(input, &data)

	if err != nil {
		printWarn("unable to load data")
		os.Exit(1)
	}

	return data
}
