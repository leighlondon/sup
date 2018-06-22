package storage

import (
	"testing"
)

func testFilenameReturnedCorrectly(t *testing.T) {
	s, err := New("tmp")
	if err != nil {
		t.Error(err)
	}
	if s.Filename() != "tmp" {
		t.Error("filename not returned properly")
	}
}
