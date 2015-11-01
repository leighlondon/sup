package main

import (
	"os"
	"testing"
)

// If the environment variable filename fails for whatever reason
// or it's not set, then there needs to be a default filename to
// fall back on.
func TestFallsBackToDefaultFilename(t *testing.T) {
	// Unset the environment variable if it's there.
	tempFileEnv := os.Getenv(fileEnv)
	if tempFileEnv != "" {
		os.Unsetenv(fileEnv)
	}
	// The filename should never be empty.
	if getFilename() == "" {
		t.Error("no default filename set")
		// Set the environment variable again if it was set initially.
		if tempFileEnv != "" {
			os.Setenv(fileEnv, tempFileEnv)
		}
	}
}

// If the environment variable is set, it should use that setting.
func TestUsesFilenameFromEnvironmentWhenSet(t *testing.T) {
	// Get the storage in the env right now.
	env := os.Getenv(fileEnv)
	// Fill it with a temporary name if it's not set, but make sure
	// to unset it when done.
	if env == "" {
		env = "TEST_SUPERB_FILENAME"
		os.Setenv(fileEnv, env)
		defer os.Unsetenv(fileEnv)
	}
	filename := getFilename()
	if filename != env {
		t.Error("doesn't use the env for filename")
	}
}
