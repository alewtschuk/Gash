package main

import (
	"gash/go-src/shell"
	"gash/go-src/shell/parser"
	"os"
)

func main() {

	const version string = "1.0.0"
	// Call ParseArgs to parse any args passed to Gash
	parser.ParseStartArgs(len(os.Args), version)
	shell.Run()
}
