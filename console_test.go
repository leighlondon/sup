package main

import "testing"

func TestNewlineNotAddedWhenPresent(t *testing.T) {
	// Don't append a newline when normalising if there's
	// already one there.
	hasNewline := "newline\n"
	newString := fixNewline(hasNewline)

	if newString != hasNewline {
		t.Error("added incorrect newline")
	}
}

func TestNewlineAddedWhenNeeded(t *testing.T) {
	// Append a newline if it's not present.
	noNewline := "no newline"
	newString := fixNewline(noNewline)

	if newString != noNewline+"\n" {
		t.Error("didn't add newline when needed")
	}
}

func TestExcessWhitespaceRemoved(t *testing.T) {
	// Too many extra lines should be removed and trimmed.
	tooMany := "too many newlines"
	newString := fixNewline(tooMany + "\n\n\n")

	if newString != tooMany+"\n" {
		t.Error("didn't remove the excess whitespace")
	}
}
