package main

import "testing"

// Don't append a newline when normalising if there's
// already one there.
func TestNewlineNotAddedWhenPresent(t *testing.T) {
	hasNewline := "newline\n"
	newString := fixNewline(hasNewline)
	if newString != hasNewline {
		t.Error("added incorrect newline")
	}
}

// Append a newline if it's not present.
func TestNewlineAddedWhenNeeded(t *testing.T) {
	noNewline := "no newline"
	newString := fixNewline(noNewline)
	if newString != noNewline+"\n" {
		t.Error("didn't add newline when needed")
	}
}

// Too many extra lines should be removed and trimmed.
func TestExcessWhitespaceRemoved(t *testing.T) {
	tooMany := "too many newlines"
	newString := fixNewline(tooMany + "\n\n\n")
	if newString != tooMany+"\n" {
		t.Error("didn't remove the excess whitespace")
	}
}
