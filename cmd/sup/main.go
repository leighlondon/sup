package main

import (
	"flag"
	"log"
	"os"

	"github.com/leighlondon/sup"
	"github.com/leighlondon/sup/pkg/config"
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

func run(store sup.Storer, opts sup.Options, out *log.Logger, args ...string) {
	if opts.Version {
		out.Println(sup.Version)
		return
	}
	if opts.Filename {
		out.Println(store.Filename())
		return
	}
	if opts.All || len(args) < 1 {
		for k, v := range store.All() {
			out.Printf("%s: %s", k, v)
		}
		return
	}
	if len(args) == 1 {
		v, err := store.Get(args[0])
		if err != nil {
			out.Printf("%s", err.Error())
			return
		}
		out.Println(v)
		return
	}
	if len(args) == 2 {
		k := args[0]
		v := args[1]
		store.Put(k, v)
		return
	}
}

func main() {

	output := log.New(os.Stdout, "", 0)

	var opts = sup.Options{}

	// Add the opts.
	flag.BoolVar(&opts.All, "a", false, "Show all of the keys and values.")
	flag.BoolVar(&opts.Delete, "d", false, "Delete the listed key.")
	flag.BoolVar(&opts.Filename, "f", false, "Show the storage file path.")
	flag.BoolVar(&opts.Version, "v", false, "Show the version.")

	// Set a prettier "usage" screen, for "-h" and "--help" opts.
	flag.Usage = func() { output.Printf("%s", usage) }

	// Parse the opts.
	flag.Parse()

	path, err := config.Filepath()
	if err != nil {
		output.Println(err)
		os.Exit(1)
	}

	args := flag.Args()
	store, err := storage.New(path)
	if err != nil {
		output.Println(err)
		os.Exit(2)
	}

	run(store, opts, output, args...)
}
