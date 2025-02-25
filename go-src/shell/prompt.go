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

func updatePrompt(updateStr string) {
	var currentPrompt string = GetPrompt("") // TODO: Adjust for envVar
	//var hostname, _ = os.Hostname()
	//var user, _ = user.Current()
	//var username string = user.Username

	before, after, _ := strings.Cut(currentPrompt, " ")
	log.Println("Before: " + before)
	log.Println("After: " + after)
	log.Println(os.Getwd())
}
