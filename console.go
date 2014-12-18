package main

import (
	"fmt"
	"os"
	"strings"
)

func printWarn(s string) {
	fmt.Fprintf(os.Stderr, fixNewline(s))
}

func fixNewline(s string) string {

	var output string

	if strings.HasSuffix(s, "\n") {
		output = s
	} else {
		output = s + "\n"
	}

	return output
}
