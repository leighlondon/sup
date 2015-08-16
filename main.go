package main

import (
	"flag"
)

const usageString = `superb ` + Version + `

Usage:
  sup [options]
  sup [options] <key> [<value>]

Options:
  -a --all        Show all of the keys and values.
  -d --delete     Delete the listed key.
  -f --file       Show the storage file path.
  -h --help       Show this screen.
  -v --version    Show the version.
`

func main() {

	// Add the flags.
	var allFlag = flag.Bool("a", false, "Show all of the keys and values.")
	var deleteFlag = flag.Bool("d", false, "Delete the listed key.")
	var filenameFlag = flag.Bool("f", false, "Show the storage file path.")
	var versionFlag = flag.Bool("v", false, "Show the version.")

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
		printWarn(VersionString)
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
