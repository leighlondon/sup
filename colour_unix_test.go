package main

import (
	"strings"
	"testing"
)

func TestColouredStingsHaveEscapeCodes(t *testing.T) {
	testString := "test"
	coloured := wrapColour("red", testString)

	if !strings.Contains(coloured, "\x1b") {
		t.Error("didn't contain valid color code")
	}

	if !strings.Contains(coloured, "\x1b[;31m") {
		t.Error("didn't contain the correct colour code")
	}
}

func TestInvalidColoursReturnTheSameString(t *testing.T) {
	testString := "test"
	coloured := wrapColour("blurple", testString)

	if strings.Contains(coloured, "\x1b") {
		t.Error("added an escape code incorrectly")
	}

	if !strings.EqualFold(coloured, testString) {
		t.Error("invalid color should be the same string")
	}
}
