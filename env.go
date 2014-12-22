package main

import (
	"os"
)

const supFile = ".superb"
const fileEnv = "SUPERB_FILE"

func getFilename() string {

	var name string

	env := os.Getenv(fileEnv)

	if env != "" {
		name = env
	} else {
		name = supFile
	}

	return name
}
