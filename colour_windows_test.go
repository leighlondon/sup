// +build windows

package main

import "testing"

func TestValidColourOnWindowsDoesNothing(t *testing.T) {
	// Attempting to wrap a string with a valid colour on Windows
	// should not change the string in any way.
	testString := "test"
	unColoured := WrapColour("red", testString)

	// The strings should be literally identical.
	if testString != unColoured {
		t.Error("valid windows coloured string wasn't identical")
	}
}

func TestInvalidColourOnWindowsDoesNothing(t *testing.T) {
	// Attempting to wrap with an invalid colour on Windows should not
	// change the string in any way.
	testString := "test"
	unColoured := WrapColour("blurple", testString)

	// The strings should be literally identical.
	if testString != unColoured {
		t.Error("invalid windows coloured string wasn't identical")
	}
}
