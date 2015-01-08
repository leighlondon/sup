// Console output helpers.
package main

import (
	"fmt"
	"os"
	"strings"
)

// Printing to standard output.
func printOut(s string) {
	fmt.Fprintf(os.Stdout, fixNewline(s))
}

// Printing to standard error.
func printWarn(s string) {
	fmt.Fprintf(os.Stderr, fixNewline(s))
}

// Normalise any leading and trailing whitespace.
func fixNewline(s string) string {
	output := strings.TrimSpace(s)
	return output + "\n"
}
