package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

const Version = `0.0.0`
const usage = `superb.

Usage:
  sup <key> <value>
  sup [options]

Options:
  -h --help	  Show this screen.
  -v --version	  Show the version.
`

func main() {

	arguments, _ := docopt.Parse(usage, nil, false, Version, false)

	if arguments["--help"] == true {
		fmt.Println(usage)
	}
}
