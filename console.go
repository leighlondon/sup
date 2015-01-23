// Console output helpers.
package main

import (
	"fmt"
	"os"
	"strings"
)

// Printing to standard output.
func printOut(s string) {
	output := fixNewline(s)
	fmt.Fprintf(os.Stdout, output)
}

// Printing to standard error.
func printWarn(s string) {
	output := fixNewline(s)
	output = wrapColour("blue", output)
	fmt.Fprintf(os.Stdout, output)
}

func printKeyValue(key, value string) {
	keyString := wrapColour("blue", key)
	valueString := wrapColour("red", value)
	printOut(keyString + " => " + valueString)
}

// Normalise any leading and trailing whitespace.
func fixNewline(s string) string {
	output := strings.TrimSpace(s)
	return output + "\n"
}
