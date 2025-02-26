package shell

// Runs the shell
func Run() {
	// Initalize reader
	initReader()
	for {
		line := readLine()
		//log.Print(line) // Uses log to avoid buffering issues

		// Skip processing line if empty
		if len(line) == 0 {
			continue
		}

		var isBuiltin bool = getCommandType(line)
		if !isBuiltin {
			execute(line)
		} else {
			handleBuiltins(line)
		}

	}
}

// Checks if the command is an internal or
// external command.
func getCommandType(line []string) bool {

	// Skip processing line if empty
	if len(line) == 0 {
		return false
	}

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
