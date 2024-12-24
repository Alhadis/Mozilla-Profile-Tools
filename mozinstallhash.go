package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bradenhilton/mozillainstallhash"
)

const usage = `
mozinstallhash
	Get the hash used to differentiate between installs of Mozilla software.

Usage:
	mozinstallhash <path> [<path> ...]

Where <path> is a string describing the parent directory of the executable,
e.g. "C:\Program Files\Mozilla Firefox", with platform specific path separators
("\" on Windows and "/" on Unix-like systems)

Example:
	mozinstallhash "C:\Program Files\Mozilla Firefox"
	308046B0AF4A39CB

	mozinstallhash "C:/Program Files/Mozilla Firefox"
	9D561FCD08DC6D55

	mozinstallhash "/usr/lib/firefox"
	4F96D1932A9F858E`

func main() {
	if len(os.Args) == 1 {
		log.Println(fmt.Errorf("error: no path specified"))
		fmt.Println(usage)
		os.Exit(1)
	}

	paths := os.Args[1:]
	for _, path := range paths {
		path = strings.TrimSuffix(path, "/")
		path = strings.TrimSuffix(path, "\\")

		hash, err := mozillainstallhash.MozillaInstallHash(path)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(hash)
	}
}
