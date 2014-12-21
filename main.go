package main

import (
	"os"

	"github.com/docopt/docopt-go"
)

func main() {

	arguments, _ := docopt.Parse(usage, nil, false, "", false)

	// The options flags are only one at a time each, and
	// should exit immediately when completed.
	if arguments["--help"] == true {
		printWarn(usage)
		os.Exit(0)
	} else if arguments["--version"] == true {
		printWarn(VersionString)
		os.Exit(0)
	}

	// Load in the data file.
	data := Load("test.json")

	// Need to use a type assertion for a string on the "key" argument.
	key, ok := arguments["<key>"].(string)
	if !ok {
		// No key provided, just give the usage screen and exit.
		printWarn(usage)
		os.Exit(0)
	}

	if value, ok := data[key]; ok {
		printOut(value)
	} else {
		printWarn("not found")
	}
}
