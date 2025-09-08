# Gash 
Gash is a minimalist Unix shell implemented in Go. It serves as a lightweight and extensible command-line interface. WIP, will be expanded.
Currently being refactored in refactor branch. 

**Project is currently in alpha state**

* Current Core Functionality:
  * Executes external commands from PATH.
  * Basic built-in commands: cd, exit, history.
  * Dynamic prompt showing user@hostname:path.
  * Readline support for command history and editing.
  * Tab completion for commands, files, and directories.

* Known Bugs:
  * cd command updates the prompt even on failure.
  * Ctrl-C in a child process can make the shell unresponsive.

TODO Roadmap :
- Fix existing bugs
- Piping
- I/O redirection
- Scripting support
- Background processes
- Configuration file
- And More

Steps to configure, build, run, and test the project.

## Building

```bash
make build
```

## Testing

```bash
make check
```

## Clean

```bash
make clean
```
