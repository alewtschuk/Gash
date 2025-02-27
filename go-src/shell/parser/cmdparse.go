package parser

import (
	"log"
	"os"
	"path"
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

	// Check for reference of home dir using "~" or "~/"
	if len(args) > 1 {
		args = checkHomeRef(args)
	}

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

// Checks if there is a home directory reference using "~"" or "~/"
// and replaces it with the path of the home directory
func checkHomeRef(args []string) []string {
	var tilde, tslash = "~", "~/"
	var homeDir, err = os.UserHomeDir()
	if err != nil {
		log.Println("Error getting home directory")
	}

	// For each argument in args at index i
	for i, arg := range args {
		//log.Println("Arg is " + arg)
		// If arg at i contains "~"" or "~/" replace refernce with home directory
		switch {
		case strings.Contains(arg, tilde):
			args[i] = path.Clean(strings.Replace(arg, tilde, homeDir, 1))
		case strings.Contains(arg, tslash):
			args[i] = path.Clean(strings.Replace(arg, tslash, homeDir, 1))
		default:
			// If not present do nothing
			continue
		}
		///log.Println("Arg is now " + arg)
	}
	//log.Println(args)

	return args
}
