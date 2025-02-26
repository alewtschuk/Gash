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
		Prompt:          GetPrompt("PROMPT"),
		HistoryFile:     "/tmp/gashcmds.tmp",
		AutoComplete:    setCompleter(),
		InterruptPrompt: "^C",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
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

// Function to filter input to avoid some issues handling signals
func filterInput(r rune) (rune, bool) {
	switch r {
	// Block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
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
func readLine() []string {

	line, err := reader.Readline()

	// Capture and ignore exit signals
	reader.CaptureExitSignal()
	if err == readline.ErrInterrupt {
		return nil
	} else if err == io.EOF {
		return nil
	} else if err != nil {
		return nil
	}

	// Get the parsed command slice from the parser
	var parsedline []string = parser.ParseCommand(line)
	if parsedline == nil {
		log.Println("Unable to parse line")
	}

	return parsedline
}
