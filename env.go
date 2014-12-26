package main

import (
	"os"
	"os/user"
	"path/filepath"
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

// Get the directory containing the storage file.
// This checks for an environment variable (SUPERB_DIR) and falls back
// to the HOME environment variable for the current user.
func getFileDirectory() string {

	// Try the environment.
	dir := os.Getenv(dirEnv)


	if dir == "" {
		// Default to the home directory for the user
		// if the env variable wasn't set.
		usr, err := user.Current()
		dir = usr.HomeDir

		if err != nil {
			printWarn("problem with home directory")
			os.Exit(1)
		}
	}

	return dir
}

// Getting the file path from the environment.
func FilePath() string {
	// Join the filepath and the filename together with the
	// OS-specific separator.
	return filepath.Join(getFileDirectory(), getFilename())
}
