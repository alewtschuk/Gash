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
	// Create map to store flag values
	var flagMap map[string]any = make(map[string]any)

	// Set and parse flags
	var versionFlag *bool = flag.Bool("v", false, "Check Gash version")
	flag.Parse()

	// Store flag values in the map
	flagMap["version"] = *versionFlag

	// Switch on argCount to handle parsing for arguments
	switch {
	case argCount == 2:
		// If version is true print gash version and exit
		if flagMap["version"].(bool) {
			fmt.Println("Gash version " + version)
			os.Exit(0)
		}
	case argCount > 2:
		log.Fatal("Multiple arguments not supported")
	default:
		break
	}
}
