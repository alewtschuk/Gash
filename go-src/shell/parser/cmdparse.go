package parser

import (
	"os"
	"strings"

	"github.com/alewtschuk/dsutils"
)

// Get a list of all executable commands in the PATH
func ExtractPathExecutatbles() []string {
	// Create path slice by splitting the PATH value at : delimiter
	var paths []string = strings.Split(os.Getenv("PATH"), ":")
	var commandList []string
	// For each path in the paths append the filename to the commands slice
	for _, path := range paths {
		files, err := os.ReadDir(path)
		if err != nil {
			continue
		}
		for _, file := range files {
			commandList = append(commandList, file.Name())
		}
	}
	// Ensure commands slice only contains unique values
	commandList = dsutils.EnsureUnique(commandList)
	return commandList
}

// Parses the command into a slice of strings to be returned
func ParseCommand(line string) []string {
	// Split line into slice of strings at whitespace
	line = strings.Trim(line, " ") // TODO: Check for all invisible unicode characters as future edge case
	return strings.Split(line, " ")
}
