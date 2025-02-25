package shell

import (
	"log"
	"os"
)

// Define builtin commands
var builtins []string = []string{"cd", "exit", "history"}

// Retrieves the needed environment variables from the OS
// for handling directories
func getDirEnvars() (string, string, string) {
	// Declare all needed vars
	// Get working directory environment variable
	var pwd string = os.Getenv("PWD")
	// Get the old working directory environment variable
	var oldpwd string = os.Getenv("OLDPWD")
	// Get the user's home directory environment variable
	var home = os.Getenv("HOME")

	log.Println("DEBUG: Current working directory: " + pwd)
	log.Println("DEBUG: Old working directory: " + oldpwd)
	log.Println("DEBUG: User home directory: " + home)

	return pwd, oldpwd, home
}

// Changes the current working directory of the shell.
// Uses the *nix system call chdir. With no arguments the users home
// directory is used as the defaut.
func changeDir(dir string) (int, error) {
	// BUG: First run does not change directory if subdir is present, second run does
	// TODO: Check if the dir passed is a subdir  of the current and handle accordingly

	var err error
	var pwd, _, home = getDirEnvars()

	// If cd recieved with no path specified change dir to home
	if dir == "cd" {
		log.Println("DEBUG: Cd passed with no args changing directory to home")
		err = os.Chdir(home)
		if err != nil {
			log.Print("DEBUG: Error changing directory ")
			log.Println(err)
			return -1, err
		}
	}

	// Attempt directory change to target dir input
	// err = os.Chdir(dir)
	// log.Println("DEBUG: Current working directory post change attempt: " + pwd)
	// // If there is an error changing the dir return error values
	// if err != nil {
	// 	log.Print("DEBUG: Error changing directory ")
	// 	log.Println(err)
	// 	return -1, err
	// }

	// Attempt directory change with user specified dir
	// if error returns use users home directory as value
	// var error = os.Chdir(dir)

	// if error != nil {
	// 	error = os.Chdir(home)
	// }
	log.Println("DEBUG: Current wokring directory after default: " + pwd + "\n")
	return 0, err
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
			var dirReturn int
			var err error
			// If cd called with non args pass "cd" to changeDir and eval accordingly
			if len(cmd) == 1 {
				changeDir(cmd[0])
			}

			// If cd called with arg pass the file path cmd[1] to changeDir and eval accordingly
			if len(cmd) > 1 {
				dirReturn, err = changeDir(cmd[1])
				//updatePrompt(cmd[1])
			}

			// If error values returned from changeDir print error message
			if err != nil || dirReturn == -1 {
				log.Println(cmd[0] + "no such file or directory: " + cmd[1])
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

// Simple error checking function.
// Only used for File reading.
func errcheck(e error) {
	if e != nil {
		panic(e)
	}
}
