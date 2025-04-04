package shell

// Declare shell struct
// Holds the InputReader wrapper for the reader,
// Executor and BuiltinsHandler.
//
// Also possibly holds the config
type Shell struct {
	io      IO
	handler Handler
	Config  *ShellConfig // Will not be a pointer (probably)
}

// Should be the only shell variable referenced
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

		// Evaluates the command
		globalShell.handler.handle(cmd)
	}
}

// Constructs the new shell struct
//
// TODO: implement and fill out actual internal functions
func NewShell() *Shell {
	// Makes assignment, assignes memory, and evaluates functions storing results in struct fields
	sh := &Shell{
		io: IO{
			Reader: newInputReader(),
			Parser: newParser(),
		},
		handler: Handler{
			Executor: newExecutor(),
			Builtins: newBuiltins(),
		},
		Config: newShellConfig(),
	}
	return sh
}

// Check and assign global shell
func GetShell() *Shell {
	if globalShell == nil {
		globalShell = NewShell()
	}
	return globalShell
}
