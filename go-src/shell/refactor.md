# shell.go
- there are many function calls tied to the shell concept that are free floating, they have no reciever. this is prime for a struct.
- shell should be only control logic, read input, parse input, route to correct handler, and handle the loop/quit/signals nothing else
    - should not decide how built-ins are implemented, execute directly, know how to parse commands
- command struct should be passed around (will need to be set up for AST node parsing)

At current the needed struct to avoid passing global functions is:
```
type Shell struct {
  Reader     InputReader    // Wraps readline stuff
  Executor   Executor       // Executes external commands
  Builtins   BuiltinHandler // Interface for built-in resolution
  Config     ShellConfig    // Startup rc file, env vars, aliases, etc.
}
```

Future Shell struct could include things like:
- Env map[string]string — local shell vars (overriding os.Environ)
- Jobs or JobTable — for background process management
- Prompt or PromptRenderer — for dynamic prompts
- History — maybe separate from Reader if you persist it
- Aliases map[string]string — for alias expansion

The Command struct should be flexible to start for long-term scalability. At a minimum:
```
type Command struct {
  Name      string   // "ls"
  Args      []string // ["-la", "/etc"]
  IsBuiltin bool     // true for "cd", "exit", etc
  Background bool    // true for "sleep 1 &"
}
```

# TODO
- define shell struct and fields
- modify Run() to be a method on Shell
- the loop should be modified to:
    - read line from Shell.Reader
    - pass through parser
    - route based on Command struct's builtin boolean value
    - pass to shell.Executor or Shell.Builtins