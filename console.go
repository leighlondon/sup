package main

import (
	"fmt"
	"os"
)

func printWarn(s string) {
	fmt.Fprintf(os.Stderr, s)
}
