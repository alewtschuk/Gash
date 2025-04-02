package shell

// Declare shell struct
// Holds the InputReader wrapper for the reader,
// Executor and BuiltinsHandler.
//
// Also possibly holds the config
type Shell struct {
	Reader   InputReader
	Parser   Parser
	Executor Executor
	Builtins BuiltinHandler
	Config   *ShellConfig // Will not be a pointer (probably)
}

// Runs the shell. Method on shell struct
func (sh *Shell) Run() {
	// Initalize reader
	sh.Reader.initReader()
	for {
		line := sh.Reader.readLine()
		//log.Print(line) // Uses log to avoid buffering issues

		// Skip processing line if empty
		if len(line) == 0 {
			continue
		}

		// cmd is the central object that moves in the system, not the raw line
		// cmd must be parsed then passed to propper handler
		cmd := sh.Parser.parse(line) // TODO: Needs to be implemented in parser
		if cmd.IsBuiltin {
			sh.Builtins.handleBuiltins(cmd)
		} else {
			sh.Executor.execute(cmd)
		}

	}
}

// Constructs the new shell struct
//
// TODO: implement and fill out actual internal functions
func NewShell() *Shell {
	sh := &Shell{}
	sh.Reader = newInputReader()
	sh.Parser = newParser()
	sh.Executor = newExecutor()
	sh.Builtins = newBuiltins()
	sh.Config = newShellConfig()
	return sh
}
