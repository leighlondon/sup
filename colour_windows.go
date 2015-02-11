// +build windows

package main

// Windows builds don't have colour printing enabled, so calling the
// colour function is essentially a no-op.
func WrapColour(colour, input string) string {
	return input
}
