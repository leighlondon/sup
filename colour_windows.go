// +build windows

package main

func WrapColour(colour, input string) string {

	// Windows builds don't have colour printing enabled, so
	// calling the colour function is essentially a no-op.
	return input
}
