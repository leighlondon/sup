// Console colouring.
package main

// Windows builds don't have colour printing enabled, so calling the
// colour function is essentially a no-op.
func wrapColour(colour, input string) string {
	return input
}
