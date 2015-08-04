package main

import (
	"flag"
)

var (
	allFlag      bool
	deleteFlag   bool
	filenameFlag bool
	versionFlag  bool
)

func init() {
	// Add the flags.
	flag.BoolVar(&allFlag, "a", false, "Show all of the keys and values.")
	flag.BoolVar(&deleteFlag, "d", false, "Delete the listed key.")
	flag.BoolVar(&filenameFlag, "f", false, "Show the storage file path.")
	flag.BoolVar(&versionFlag, "v", false, "Show the version.")

	// Parse the flags.
	flag.Parse()
}

func main() {

	// Get the filename from the environment.
	filename := FilePath()

	// The options flags are only one at a time each, and
	// should exit immediately when completed.
	if versionFlag {
		printWarn(VersionString)
		return
	} else if filenameFlag {
		printOut(filename)
		return
	}

	// Load in the data file.
	data := Load(filename)

	// Now that the data has been loaded it can check for the
	// data related options.
	if allFlag {
		for key, value := range data {
			printKeyValue(key, value)
		}
		return
	}

	// No key provided, just give the usage screen and exit.
	if flag.NArg() == 0 {
		printWarn(usage)
		return
	}

	// The key is the first argument after the flags.
	key := flag.Arg(0)

	// Check now for the delete flag.
	if deleteFlag {
		delete(data, key)
		Save(filename, data)
		return
	}

	// If the "value" argument field is present just save the key-value
	// pair into the storage.
	if flag.NArg() == 2 {
		value := flag.Arg(1)
		data[key] = value
		Save(filename, data)
	}

	// If a key-value pair is present, print it.
	if value, ok := data[key]; ok {
		printKeyValue(key, value)
	} else {
		printWarn("not found")
	}
}
