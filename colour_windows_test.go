// Testing the Windows colour printing features.

package main

import (
	"testing"
)

func TestValidColourOnWindowsDoesNothing(t *testing.T) {
	testString := "test"
	unColoured := wrapColour("red", testString)

	if testString != unColoured {
		t.Error("windows coloured string wasn't identical")
	}
}
