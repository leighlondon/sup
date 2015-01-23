// Testing the Windows colour printing features.

package main

import (
	"testing"
)

func TestValidColourOnWindowsDoesNothing(t *testing.T) {
	// Attempting to wrap a string with a valid colour on Windows
	// should not change the string in any way.
	testString := "test"
	unColoured := wrapColour("red", testString)

	// The strings should be literally identical.
	if testString != unColoured {
		t.Error("windows coloured string wasn't identical")
	}
}

func TestInvalidColourOnWindowsDoesNothing(t *testing.T) {
	// Attempting to wrap with an invalid colour on Windows should not
	// change the string in any way.
	testString := "test"
	unColoured := wrapColour("blurple", testString)

	// The strings should be literally identical.
	if testString != unColoured {
		t.Error("windows colour code wasn't base string")
	}
}
