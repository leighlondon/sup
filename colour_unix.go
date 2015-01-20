// Console colouring.
package main

// Coloured escape codes.
var colourMap = map[string]string{
	"red":     "\x1b[;31m",
	"green":   "\x1b[;32m",
	"yellow":  "\x1b[;33m",
	"blue":    "\x1b[;34m",
	"magenta": "\x1b[;35m",
	"cyan":    "\x1b[;36m",
	"white":   "\x1b[;37m",
	"reset":   "\x1b[0m",
}

// Wrap a provided string with ANSI colour codes for terminal output.
func wrapColour(colour, input string) string {
	return colourMap[colour] + input + colourMap["reset"]
}
