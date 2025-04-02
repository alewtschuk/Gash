package parser

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Parses cli arguments passed when calling gash from the command line.
// Currently only has -v version flag.
// Uses a map to be able to dynamically expand args in the future.
func ParseStartArgs(argCount int, version string) {

	// Set and parse flags
	var versionFlag *bool = flag.Bool("v", false, "Check Gash version")
	flag.Parse()

	// Switch on argCount to handle parsing for arguments
	switch {
	case argCount == 2:
		// If version is true print gash version and exit
		if *versionFlag {
			fmt.Println("Gash version " + version)
			os.Exit(0)
		}
	case argCount > 2:
		log.Fatal("Multiple arguments not supported")
	default:
		break
	}
}
