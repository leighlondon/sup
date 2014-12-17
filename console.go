package main

import (
	"fmt"
	"os"
	"strings"
)

func printWarn(s string) {

	var output string

	if strings.HasSuffix(s, "\n") {
		output = s
	} else {
		output = s + "\n"
	}

	fmt.Fprintf(os.Stderr, output)
}
