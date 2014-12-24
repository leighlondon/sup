package main

import (
	"os"
)

// The default filename to fall back to.
const supFile = ".superb"

// The environment variable for the filename.
const fileEnv = "SUPERB_FILE"

// The environment variable for the directory.
const dirEnv = "SUPERB_DIR"

// Get the filename to be used for the storage.
// This will check for an environment variable (SUPERB_FILE) and falls
// back to a default filename if this variable is not set.
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

func getFileDirectory() string {

	var dir string

	home := os.Getenv("HOME")
	env := os.Getenv(dirEnv)

	if env != "" {
		dir = env
	} else {
		dir = home
	}

	return dir
}
