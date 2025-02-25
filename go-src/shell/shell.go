package shell

import (
	"fmt"

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
		fmt.Println("DEBUG: readLine() returned ->" + line) // ✅ Step 5 Debug
	}
}

// Converts line read from the user into a format
// uses os.StartProcess(). The number of args will
// be limited to the number of max arguments loaded
// from sysconf.
