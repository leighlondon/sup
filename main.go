package main

import (
	"github.com/docopt/docopt-go"
)

func main() {

	arguments, _ := docopt.Parse(usage, nil, false, "", false)

	if arguments["--help"] == true {
		printWarn(usage)
	} else if arguments["--version"] == true {
		printWarn(Version)
	}
}
