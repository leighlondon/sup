package main

import (
	"flag"
	"fmt"

	"github.com/leighlondon/sup"
	"github.com/leighlondon/sup/pkg/storage"
)

const usage = `superb ` + sup.Version + `

Usage:
    sup [options]
    sup [options] <key> [<value>]

Options:
    -a     Show all of the keys and values.
    -d     Delete the listed key.
    -f     Show the storage file path.
    -h     Show this screen.
    -v     Show the version.
`

var (
	all      bool
	delete   bool
	filename bool
	version  bool
)

func run(store sup.Storer, args ...string) {
	if version {
		fmt.Println(sup.Version)
		return
	}
	if filename {
		fmt.Println(store.Filename())
		return
	}
	if all {
		for k, v := range store.All() {
			fmt.Printf("%s: %s", k, v)
		}
		return
	}
	if len(args) < 1 {
		fmt.Println("invalid")
		return
	}
}

func main() {

	// Add the flags.
	flag.BoolVar(&all, "a", false, "Show all of the keys and values.")
	flag.BoolVar(&delete, "d", false, "Delete the listed key.")
	flag.BoolVar(&filename, "f", false, "Show the storage file path.")
	flag.BoolVar(&version, "v", false, "Show the version.")

	// Set a prettier "usage" screen, for "-h" and "--help" flags.
	flag.Usage = func() { fmt.Printf("%s", usage) }

	// Parse the flags.
	flag.Parse()

	store := storage.NewInMemoryStorage()
	args := flag.Args()

	run(store, args...)
}
