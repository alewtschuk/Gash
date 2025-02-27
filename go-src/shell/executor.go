package shell

import (
	"log"
	"os"
	"os/exec"
)

func execute(args []string) {
	//log.Println("DEBUG: Function execute called. Recieved ", args)

	cmd := exec.Command(args[0], args[1:]...)
	setCmdStds(cmd)
	//log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	if err != nil {
		log.Printf("Command %s unable to be ran...\n", args[0])
		//log.Printf("Error is: %v\n", err)
	}
}

func setCmdStds(cmd *exec.Cmd) {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
