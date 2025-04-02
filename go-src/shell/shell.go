package shell

// Declare shell struct
// Holds the InputReader wrapper for the reader,
// Executor and BuiltinsHandler.
//
// Also possibly holds the config
type Shell struct {
	// Data for io
	io struct {
		Reader InputReader
		Parser Parser
	}
	// Data for handlers
	handler struct {
		Executor Executor
		Builtins BuiltinHandler
	}
	Config *ShellConfig // Will not be a pointer (probably)
}

var globalShell *Shell = nil

// Runs the shell. Method on shell struct
func (sh *Shell) Run() {
	// Initalize reader
	sh.io.Reader.initReader()
	for {
		line := sh.io.Reader.readLine()
		//log.Print(line) // Uses log to avoid buffering issues

		// Skip processing line if empty
		if len(line) == 0 {
			continue
		}

		// cmd is the central object that moves in the system, not the raw line
		// cmd must be parsed then passed to propper handler
		cmd := sh.io.Parser.parse(line) // TODO: Needs to be implemented in parser
		if cmd.IsBuiltin {
			sh.handler.Builtins.handleBuiltins(cmd)
		} else {
			sh.handler.Executor.execute(cmd)
		}

	}
}

// Constructs the new shell struct
//
// TODO: implement and fill out actual internal functions
func NewShell() *Shell {
	sh := &Shell{
		io: struct {
			Reader InputReader
			Parser Parser
		},
		handler: struct {
			Executor Executor
			Builtins BuiltinHandler
		},
		Config: newShellConfig(),
	}
	sh.Reader = newInputReader()
	sh.Parser = newParser()
	sh.Executor = newExecutor()
	sh.Builtins = newBuiltins()
	sh.Config = newShellConfig()
	globalShell = sh // Sets global shell to newshell
	return sh
}
