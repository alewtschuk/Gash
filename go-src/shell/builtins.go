package shell

import (
	"log"
	"os"
)

// Define builtin commands
var builtins []string = []string{"cd", "exit", "history"}

// Changes the current working directory of the shell.
// Uses the *nix system call chdir. With no arguments the users home
// directory is used as the defaut.
func changeDir(dir string) (int, error) {
	//Get the user's home directory
	var homeDir, error = os.UserHomeDir()
	if error != nil {
		return -1, error
	}
	// Attempt directory change with user specified dir
	// if error returns use users home directory as value
	error = os.Chdir(dir)
	wd, _ := os.Getwd()
	log.Println("Current wokring directory: " + wd)
	if error != nil {
		error = os.Chdir(homeDir)
	}
	return 0, error
}

// Returns the builtin command values
func getBuiltins() []string {
	return builtins
}

func handleBuiltins(cmd []string) {
	if len(cmd) > 1 && cmd[0] != "cd" {
		log.Println("To many arguments passed. Command not recognized")
		return
	} else {
		switch cmd[0] {
		case "cd":
			var dirReturn, err = changeDir(cmd[1])
			if err != nil || dirReturn == -1 {
				log.Println(cmd[0] + "no such file or directory: " + cmd[1])
			}

			if len(cmd) > 1 {
				//updatePrompt(cmd[1])
			}
		case "exit":
			os.Exit(0)
		case "history":
			var file, err = os.ReadFile("/tmp/gashcmds.tmp")
			errcheck(err)
			log.Printf("%s", file)
		}
	}
}

func errcheck(e error) {
	if e != nil {
		panic(e)
	}
}
