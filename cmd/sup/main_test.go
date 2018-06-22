package main

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"github.com/leighlondon/sup"
	"github.com/leighlondon/sup/pkg/storage"
)

func TestUsageIsDefined(t *testing.T) {
	if usage == "" {
		t.Error("usage is required")
	}
}

func TestRunWithVersion(t *testing.T) {
	var b bytes.Buffer
	store := storage.NewInMemoryStorage("")
	opts := sup.Options{Version: true}
	logger := log.New(&b, "", 0)

	run(store, opts, logger)

	s := b.String()

	if !strings.Contains(s, sup.Version) {
		t.Errorf("expected '%s', got '%s'", sup.Version, s)
	}
}
