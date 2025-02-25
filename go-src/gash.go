package main

import (
	"gash/go-src/shell"
	"gash/go-src/shell/parser"
	"log"
	"os"
)

func main() {

	const version string = "1.0.0"
	// Sets stdout to be unbuffered. This is needed to propperly handle readline returns
	log.SetOutput(os.Stdout)

	// Call ParseArgs to parse any args passed to Gash
	parser.ParseStartArgs(len(os.Args), version)
	shell.Run()
}
