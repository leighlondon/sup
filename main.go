package main

import (
	"flag"
)

const usageString = `superb ` + Version + `

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

// Storage is the usage contract that the backend agrees to, and is a
// generic interface and can swap in backends so long as they support
// a simple key-value model.
type Storage interface {
	LoadData(string) map[string]string
	SaveData(string, map[string]string)
}

func main() {
	// Add the flags.
	allFlag := flag.Bool("a", false, "Show all of the keys and values.")
	deleteFlag := flag.Bool("d", false, "Delete the listed key.")
	filenameFlag := flag.Bool("f", false, "Show the storage file path.")
	versionFlag := flag.Bool("v", false, "Show the version.")
	// Set a prettier "usage" screen, for "-h" and "--help" flags.
	flag.Usage = func() {
		printOut(usageString)
		return
	}
	// Parse the flags.
	flag.Parse()
	// Get the filename from the environment.
	filename := FilePathFromEnv()
	// The options flags are only one at a time each, and
	// should exit immediately when completed.
	if *versionFlag {
		printOut(VersionString)
		return
	} else if *filenameFlag {
		printOut(filename)
		return
	}
	// Load in the data file.
	data := LoadData(filename)
	// Now that the data has been loaded it can check for the
	// data related options.
	if *allFlag {
		for key, value := range data {
			printKeyValue(key, value)
		}
		return
	}
	// No key provided, just give the usage screen and exit.
	if flag.NArg() == 0 {
		flag.Usage()
		return
	}
	// The key is the first argument after the flags.
	key := flag.Arg(0)
	// Check now for the delete flag.
	if *deleteFlag {
		delete(data, key)
		SaveData(filename, data)
		return
	}
	// If the "value" argument field is present just save the key-value
	// pair into the storage.
	if flag.NArg() == 2 {
		value := flag.Arg(1)
		data[key] = value
		SaveData(filename, data)
	}
	// If a key-value pair is present, print it.
	if value, ok := data[key]; ok {
		printKeyValue(key, value)
	} else {
		printWarn("not found")
	}
}
