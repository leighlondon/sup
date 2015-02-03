// +build !windows

// Console colouring.
package main

// Coloured escape codes.
var colourMap = map[string]string{
	"black":   "\x1b[;30m",
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

	// Gate the colour coding by checking the environment.
	if colourSettings == "colour" || colourSettings == "color" {
		// Wrap in the colour it provided.
		return wrapColourInternal(colour, input)
	} else if colourSettings == "clown" {
		// Wrap in "clown" codes, and discard the requested colour.
		return clownify(input)
	}

	// Fall back to return the original string unmodified.
	return input
}

func wrapColourInternal(colour, input string) string {

	// Wrap in the real colour code if it's in the known colours.
	if _, ok := colourMap[colour]; ok {
		return colourMap[colour] + input + colourMap["reset"]
	}

	// Fall back to return the original string unmodified.
	return input
}

// Clowns. Lots of clowns.
func clownify(input string) string {

	var result, temp string

	// Loop the entire string and apply a colour based on the index.
	// The four colours have been selected for the best result as:
	//   red, yellow, magenta, blue
	for i, char := range input {

		// Convert rune to string for the colouring function.
		temp = string(char)

		switch i % 4 {
		case 0:
			result += wrapColourInternal("red", temp)
		case 1:
			result += wrapColourInternal("yellow", temp)
		case 2:
			result += wrapColourInternal("magenta", temp)
		case 3:
			result += wrapColourInternal("blue", temp)
		}
	}

	return result
}
