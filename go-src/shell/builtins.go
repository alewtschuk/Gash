package shell

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Define builtin commands
var builtins []string = []string{"cd", "exit", "history"}

// Retrieves the needed environment variables from the OS
// for handling directories
func getDirEnvars() (string, string) {
	// Declare all needed vars
	// Get working directory environment variable
	var pwd string = os.Getenv("PWD")
	// Get the user's home directory environment variable
	var home = os.Getenv("HOME")

	//log.Println("DEBUG: Current working directory: " + pwd)
	//log.Println("DEBUG: User home directory: " + home + "\n")

	return pwd, home
}

func updateDirEnvars(nwd string, owd string) {
	var err error = os.Setenv("OLDPWD", owd)
	err = os.Setenv("PWD", nwd)
	if err != nil {
		log.Fatal("Error updating environment variables. Exiting...")
	}
	//log.Println("DEBUG: OLDPWD is now: " + owd)
	//log.Println("DEBUG: PWD is now: " + nwd + "\n")

}

// Changes the current working directory of the shell.
// Uses the *nix system call chdir. With no arguments the users home
// directory is used as the defaut.
func changeDir(dir string) (int, error) {

	var err error
	var pwd, home = getDirEnvars()

	// If cd recieved with no path specified change dir to home
	if dir == "cd" {
		//log.Println("DEBUG: Cd passed with no args changing directory to home")
		// Change dir to home and update env
		err = os.Chdir(home)
		updateDirEnvars(home, pwd)
		if err != nil {
			log.Print("Error changing directory")
			log.Println(err)
			return -1, err
		}
	} else {
		//log.Println("DEBUG: Changing directory to: " + dir + "\n")
		// Create nwd to be current working directory/requested directory
		var nwd string = path.Clean(filepath.Join(pwd, dir))
		//log.Println("DEBUG: requested directory is: " + dir)
		//log.Println("DEBUG: working directory is: " + pwd)
		//log.Println("DEBUG: nwd will be: " + pwd + " plus " + dir)
		//log.Println("DEBUG: nwd is: " + nwd)

		// If the requested directory contains the current working directory
		if strings.Contains(dir, pwd) {
			// Remove the overlap by replacing the working directory in the requested directory with itself to avoid duplication
			nwd = strings.Replace(dir, pwd, pwd, 1)
			//log.Println("DEBUG: nwd is now: " + nwd)
		}
		err = os.Chdir(dir)
		updateDirEnvars(nwd, pwd)
		if err != nil {
			log.Print("Error changing directory ")
			log.Println(err)
			return -1, err
		}
	}

	//pwd, _ = getDirEnvars()
	//log.Println("DEBUG: Current working directory after default: " + pwd + "\n")
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
				//log.Println("DEBUG: cd called with argument " + cmd[1])
				dirReturn, err = changeDir(cmd[1])
				//updatePrompt(cmd[1])
			}

			// If error values returned from changeDir print error message
			if err != nil || dirReturn == -1 {
				log.Println(cmd[0] + " no such file or directory: " + cmd[1])
			}
			// Set updated prompt
			updatePrompt()
			GetPrompt("PROMPT")

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
