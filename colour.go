package main

// WrapColour wraps a provided string with ANSI colour codes, if colour
// printing is possible on the platform.
func WrapColour(colour, input string) string {
	return wrapColour(colour, input)
}
