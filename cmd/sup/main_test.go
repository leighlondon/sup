package main

import (
	"testing"
)

func TestUsageIsDefined(t *testing.T) {
	if usage == "" {
		t.Error("usage is required")
	}
}
