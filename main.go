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
}
