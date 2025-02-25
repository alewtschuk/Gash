package shell

import (
	"fmt"
	"gash/go-src/shell/parser"
	"io"
	"log"
	"os"

	"github.com/chzyer/readline"
)

// Declare reader instance, initalized once
var reader *readline.Instance

func initReader() {
	var err error

	reader, err = readline.NewEx(&readline.Config{
		Prompt:          GetPrompt("gash >"),
		HistoryFile:     "/tmp/gashcmds.tmp",
		AutoComplete:    setCompleter(),
		InterruptPrompt: "^C",
	})

	if err != nil {
		log.Fatal("DEBUG: Failed to initialize reader:", err) // ✅ Step 7 Debug
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
	if err == readline.ErrInterrupt {
		fmt.Println("Received interrupt, exiting...")
		os.Exit(0)
	} else if err == io.EOF {
		fmt.Println("Received EOF, exiting...")
		os.Exit(1)
	} else if err != nil {
		fmt.Println("Readline error:", err) // ✅ Step 2 Debug
		return ""
	}

	return line
}
