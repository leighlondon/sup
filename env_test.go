package main

import (
	"os"
	"testing"
)

func TestFallsBackToDefaultFilename(t *testing.T) {
	// If the environment variable filename fails for whatever reason
	// or it's not set, then there needs to be a default filename to
	// fall back on.

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
