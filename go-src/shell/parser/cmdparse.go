package parser

import (
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/alewtschuk/dsutils"
)

/*
#include <unistd.h>
long getArgMax() {
	return sysconf(_SC_ARG_MAX);
}
*/
import "C"

// CGo wrapper function to get ARG_MAX from sysconf()
func getArgMax() int {
	return int(C.getArgMax())
}

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
	// Trim all leading and trailing whitespace
	line = strings.TrimSpace(line) // TODO: Check for all invisible unicode characters as future edge case

	// Split line into slice of strings at whitespaces
	var args []string = strings.Fields(line)
	// Declare max arguments byte length, pulls from sysconf using CGo
	var argMax int = getArgMax()
	var argLen int
	// For each argument get the length of the bytes of the argument and the total byte length of all args
	for _, arg := range args {
		argLen += len(arg) + 1
	}

	if argLen > argMax {
		log.Println("Arguments passed to " + args[0] + " exceed max arguments allowed by " + runtime.GOOS)
		return nil
	}
	return args
}
