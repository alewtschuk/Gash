package shell

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
)

// Sets the shell prompt. This will attempt to load the prompt
// from the requested environment variable. If the
// variable is not set a default prompt of "gash>" is returned.
func GetPrompt(envVar string) string {
	// Get User
	user, err := user.Current()
	if err != nil {
		log.Fatal("Error fetching user value. Exiting...")
	}

	// Get hostname
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Error fetching domain. Exiting...")
	}

	fpwd := getFormatedPwd()

	//Check if the prompt for the envVar key exists
	prompt, exists := os.LookupEnv("PROMPT")
	if !exists {
		//Set default value if there is no prompt
		prompt = fmt.Sprintf("%v@%v%s gash > ", user.Username, hostname, fpwd)
		// Set prompt environment variable
		os.Setenv("PROMPT", prompt)
	}

	os.Setenv("PROMPT", prompt)
	return prompt
}

func updatePrompt() {
	// Get formatted working directory
	fpwd := getFormatedPwd()
	//log.Println("FPWD is: " + fpwd)
	// Get PROMPT environment variable
	prompt := os.Getenv("PROMPT")
	//log.Println("PROMPT is: " + prompt)

	// Cut working directory at the "~" to substring after"
	splitPrompt := strings.Split(prompt, "~")

	// Set PROMPT environment variable to
	os.Setenv("PROMPT", fmt.Sprintf("%s%s gash > ", splitPrompt[0], fpwd))

	// Dynamically refresh the reader prompt
	if reader != nil {
		reader.SetPrompt(os.Getenv("PROMPT"))
		reader.Refresh()
	}
}

// Returns the formatted working directory.
// Replaces $HOME with a ~
func getFormatedPwd() string {
	// Get the current working directory
	pwd := os.Getenv("PWD")
	//log.Println("DEBUG: PWD in GetPrompt is: " + pwd)
	// Get the user home directory
	home := os.Getenv("HOME")

	return strings.Replace(pwd, home, "~", 1)
}
