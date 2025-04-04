package shell

import "fmt"

type Handler struct {
	Executor Executor
	Builtins BuiltinHandler
}

// Determines if command is built in
// handles accordingly
func (h *Handler) handle(cmd Command) {
	if cmd.IsBuiltin {
		globalShell.handler.Builtins.handleBuiltins(cmd)
	} else if cmd.IsBackground {
		fmt.Printf("%+v", cmd)
	} else {
		globalShell.handler.Executor.execute(cmd)
	}
}
