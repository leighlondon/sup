package config

import (
	"os"
	"os/user"
	"path/filepath"
)

const (
	// The default filename to fall back to.
	defaultFilename = ".superb"
	// The environment variable for the filename.
	fileEnvVar = "SUPERB_FILE"
	// The environment variable for the directory.
	dirEnvVar = "SUPERB_DIR"
)

func getFilename() string {
	name := os.Getenv(fileEnvVar)
	// Revert to the default name if the environment variable is empty.
	if name == "" {
		name = defaultFilename
	}
	return name
}

func getFileDirectory() (string, error) {
	// Try the environment.
	dir := os.Getenv(dirEnvVar)
	// Default to the home directory for the user
	// if the env variable wasn't set.
	usr, err := user.Current()
	if dir == "" {
		dir = usr.HomeDir
	}
	return dir, err
}

// Filepath returns the filepath from the environment.
func Filepath() (string, error) {
	// Join the filepath and the filename together with the
	// OS-specific separator.
	dir, err := getFileDirectory()
	file := getFilename()
	return filepath.Join(dir, file), err
}
