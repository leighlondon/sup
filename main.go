package main

import (
	"github.com/docopt/docopt-go"
)

func main() {

	arguments, _ := docopt.Parse(usage, nil, false, Version, false)

	if arguments["--help"] == true {
		printWarn(usage)
	}
}
