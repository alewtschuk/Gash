package shell

// import (

// )

type Handler struct {
	Executor Executor
	Builtins BuiltinHandler
}

func (h *Handler) handle(cmd Command) {
	if cmd.IsBuiltin {
		globalShell.handler.Builtins.handleBuiltins(cmd)
	} else {
		globalShell.handler.Executor.execute(cmd)
	}
}
