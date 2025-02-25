package shell

import (
	"log"

	"golang.org/x/sys/unix"
)

// /*
// #include <unistd.h>
// long getArgMax() {
// 	return sysconf(_SC_ARG_MAX);
// }
// */
// import "C"

// Define the struct for the shell
type shell struct {
	shell_is_interactive int
	shell_pgid           int
	shell_tmodes         unix.Termios
	shell_terminal       int
	prompt               string
}

// func InitShell() shell{
// 	return &shell{

// 	}
// }

func Run() {
	initReader()
	for {
		line := readLine()
		log.Print(line) // Uses log to avoid buffering issues
		var isBuiltin bool = getCommandType(line)
		if !isBuiltin {

		} else {
			handleBuiltins(line)
		}

	}
}

// Converts line read from the user into a format
// uses os.StartProcess(). The number of args will
// be limited to the number of max arguments loaded
// from sysconf.

func getCommandType(line []string) bool {
	// Get the list of built-in commands
	var builtins []string = getBuiltins()
	// For every builtin command check if the first arg of line matches
	for _, builtin := range builtins {
		if line[0] == builtin {
			log.Println("Command is built-in")
			return true
		}
	}
	log.Println("Command is not built-in")
	return false
}
