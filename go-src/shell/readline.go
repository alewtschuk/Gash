package shell

import (
	"gash/go-src/shell/parser"
	"io"
	"log"
	"os"

	"github.com/chzyer/readline"
)

// Declare reader instance, initalized once
var reader *readline.Instance

// Initalizes the reader
func initReader() {
	var err error

	//Log settings
	// Ensures logs print immediately
	log.SetOutput(os.Stderr)
	// Removes timestamp
	log.SetFlags(0)

	reader, err = readline.NewEx(&readline.Config{
		Prompt:          GetPrompt("gash >"),
		HistoryFile:     "/tmp/gashcmds.tmp",
		AutoComplete:    setCompleter(),
		InterruptPrompt: "^C",
	})
	if err != nil {
		log.Fatal("Failed to initialize reader")
	}
}

// Function constructor - constructs new function for listing given directory
func listFiles(path string) func(string) []string {
	return func(line string) []string {
		var names []string = make([]string, 0)
		files, _ := os.ReadDir(path)
		for _, f := range files {
			names = append(names, f.Name())
		}
		return names
	}
}

// Sets the completer to contain all internal and external commands
func setCompleter() *readline.PrefixCompleter {
	// Fetch external commands dynamically
	var externalCmds []string = parser.ExtractPathExecutatbles()

	// Fetch internal commands
	var internalCmds []string = getBuiltins()

	var rootItems []readline.PrefixCompleterInterface

	// Add built-in commands
	for _, cmd := range internalCmds {
		rootItems = append(rootItems, readline.PcItem(cmd))
	}

	// Add external commands from PATH
	for _, cmd := range externalCmds {
		rootItems = append(rootItems, readline.PcItem(cmd))
	}

	// Add file/dirrectory completion
	rootItems = append(rootItems, readline.PcItemDynamic(listFiles(".")))

	return readline.NewPrefixCompleter(rootItems...)
}

// Sets up reader and returns the line read and error if present
func readLine() string {

	line, err := reader.Readline()
	// Handles gracefull reader exit
	reader.CaptureExitSignal()
	if err == readline.ErrInterrupt {
		log.Println("Received interrupt, exiting...")
		os.Exit(0)
	} else if err == io.EOF {
		log.Println("Received EOF, exiting...")
		os.Exit(1)
	} else if err != nil {
		log.Println("Readline error:", err)
		return ""
	}

	return line
}
