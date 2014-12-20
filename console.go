package main

import (
	"fmt"
	"os"
	"strings"
)

func printOut(s string) {
	fmt.Fprintf(os.Stdout, fixNewline(s))
}

func printWarn(s string) {
	fmt.Fprintf(os.Stderr, fixNewline(s))
}

func fixNewline(s string) string {
	output := strings.TrimSpace(s)
	return output + "\n"
}
