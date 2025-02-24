package main

import (
	"flag"
	"fmt"
	"gash/go-src/shell"
	"os"
)

func main() {
	//Create map to store flag values
	var flagMap map[string]any = make(map[string]any)

	//Set and parse flags
	var versionFlag *bool = flag.Bool("v", false, "Check Gash version")
	flag.Parse()

	//Store flag values in the map
	flagMap["version"] = *versionFlag
	// Call ParseArgs to parse any args passed to Gash
	shell.ParseArgs(len(os.Args), flagMap)
	envVar := "SHELL_PROMPT"
	prompt := shell.GetPrompt(envVar)
	fmt.Print(prompt)
}
