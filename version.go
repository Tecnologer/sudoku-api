package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	versionFlag = flag.Bool("version", false, "returns the version of the build")
	version     string
	minversion  string
)

func init() {
	flag.Parse()
}

func checkVersion() {
	if len(os.Args) < 2 {
		return
	}

	if *versionFlag || argsConstainsVersion() {
		printVersion()
	}
}

func printVersion() {
	fmt.Printf("%s%s\n", version, minversion)
	os.Exit(0)
}

func argsConstainsVersion() bool {
	for _, a := range os.Args {
		if a == "version" || a == "--version" {
			return true
		}
	}

	return false
}
