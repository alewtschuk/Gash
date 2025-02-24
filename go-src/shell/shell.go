package shell

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"golang.org/x/sys/unix"
)

/*
#include <unistd.h>
long getArgMax() {
	return sysconf(_SC_ARG_MAX);
}
*/
import "C"

const version string = "1.0.0"

// Define the struct for the shell
type shell struct {
	shell_is_interactive int
	shell_pgid           int
	shell_tmodes         unix.Termios
	shell_terminal       int
	prompt               string
}

// Sets the shell prompt. This will attempt to load the prompt
// from the requested environment variable. If the
// variable is not set a default prompt of "gash>" is returned.
func GetPrompt(envVar string) string {
	//Check if the prompt for the envVar key exists
	prompt, exists := os.LookupEnv(envVar)
	user, err := user.Current()
	if err != nil {
		log.Fatal("Error fetching user value. Exiting...")
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Error fetching domain. Exiting...")
	}
	//Return default value if there is no prompt
	if !exists {
		return fmt.Sprintf("%v@%v gash >", user.Username, hostname)
	}
	return prompt
}

// Changes the current working directory of the shell.
// Uses the *nix system call chdir. With no arguments the users home
// directory is used as the defaut.
func ChangeDir(dir string) (int, error) {
	//Get the user's home directory
	var homeDir, error = os.UserHomeDir()
	if error != nil {
		return -1, error
	}
	// Attempt directory change with user specified dir
	// if error returns use users home directory as value
	error = os.Chdir(dir)
	if error != nil {
		error = os.Chdir(homeDir)
	}
	return 0, error
}

// Converts line read from the user into a format
// uses os.StartProcess(). The number of args will
// be limited to the number of max arguments loaded
// from sysconf.
// func CmdParse(line string) string {
// 	var maxArgs int = int(C.getArgMax())
// 	var funcMap map[string]func() = make(map[string]func())

// 	return fmt.Sprintln(maxArgs)
// }

func ParseArgs(argCount int, args map[string]any) {
	// Switch on argCount to handle parsing for arguments
	switch {
	case argCount == 2:
		if args["version"].(bool) {
			fmt.Println("Gash version " + version)
			os.Exit(0)
		}
	case argCount > 2:
		log.Fatal("Multiple arguments not supported")
	default:
		//TODO: Implement default call to gash if no args specified
		// fmt.Println("No arg specified")
	}
}
