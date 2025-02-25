package shell

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

// Sets the shell prompt. This will attempt to load the prompt
// from the requested environment variable. If the
// variable is not set a default prompt of "gash>" is returned.
func GetPrompt(envVar string) string {
	//Check if the prompt for the envVar key exists
	prompt, exists := os.LookupEnv(envVar)
	user, err := user.Current()
	if err != nil {
		log.Fatal("Error fetching user value. Exiting...")
	}

	// Get hostname
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Error fetching domain. Exiting...")
	}

	//Return default value if there is no prompt
	if !exists {
		return fmt.Sprintf("%v@%v gash > ", user.Username, hostname)
	}
	return prompt
}
