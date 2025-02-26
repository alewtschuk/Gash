package shell

func Run() {
	initReader()
	for {
		line := readLine()
		//log.Print(line) // Uses log to avoid buffering issues
		var isBuiltin bool = getCommandType(line)
		if !isBuiltin {
			execute(line)
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
			//log.Println("DEBUG: Command is built-in ✅\n")
			return true
		}
	}
	//log.Println("DEBUG: Command is not built-in ❌\n")
	return false
}
