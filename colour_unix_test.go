// +build !windows

package main

import (
	"strings"
	"testing"
)

func TestColouredStingsHaveEscapeCodes(t *testing.T) {
	// Attempting to wrap with a valid colour should wrap the string
	// in a valid ANSI escape code.
	testString := "test"
	coloured := WrapColour("red", testString)

	if !strings.Contains(coloured, "\x1b") {
		t.Error("didn't contain valid color code")
	}
}

func TestInvalidColoursReturnTheSameString(t *testing.T) {
	// Attempting to wrap with an invalid colour should just return
	// the original string.
	testString := "test"
	coloured := WrapColour("blurple", testString)

	if strings.Contains(coloured, "\x1b") {
		t.Error("added an escape code incorrectly")
	}

	if coloured != testString {
		t.Error("invalid color didn't return identical string")
	}
}
