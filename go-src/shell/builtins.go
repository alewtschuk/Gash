package shell

import "os"

var builtins []string = []string{"cd", "exit", "history"}

// Changes the current working directory of the shell.
// Uses the *nix system call chdir. With no arguments the users home
// directory is used as the defaut.
func ChangeDir(dir string) (int, error) {
	//Get the user's home directory
	var homeDir, error = os.UserHomeDir()
	if error != nil {
		return -1, error
	}
	// Attempt directory change with user specified dir
	// if error returns use users home directory as value
	error = os.Chdir(dir)
	if error != nil {
		error = os.Chdir(homeDir)
	}
	return 0, error
}

// Returns the builtin command values
func getBuiltins() []string {
	return builtins
}
