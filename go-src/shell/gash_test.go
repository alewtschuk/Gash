package shell

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func cmdParse(input string) []string {
	return strings.Fields(input)
}

func TestCmdParseBasic(t *testing.T) {
	input := "foo -v"
	expected := []string{"foo", "-v"}

	actual := cmdParse(input)

	if len(actual) != len(expected) {
		t.Errorf("Expected %d arguments, got %d", len(expected), len(actual))
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("Expected '%s', got '%s'", expected[i], actual[i])
		}
	}
}

func TestCmdParseMultipleArgs(t *testing.T) {
	input := "ls -a -l"
	expected := []string{"ls", "-a", "-l"}

	actual := cmdParse(input)

	if len(actual) != len(expected) {
		t.Errorf("Expected %d arguments, got %d", len(expected), len(actual))
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("Expected '%s', got '%s'", expected[i], actual[i])
		}
	}
}

func TestCmdParseEmptyInput(t *testing.T) {
	input := ""
	expected := []string{}

	actual := cmdParse(input)

	if len(actual) != len(expected) {
		t.Errorf("Expected empty slice, got %v", actual)
	}
}

func TestCmdParseExtraSpaces(t *testing.T) {
	input := "  ls    -l   -a  "
	expected := []string{"ls", "-l", "-a"}

	actual := cmdParse(input)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCmdParseSpecialCharacters(t *testing.T) {
	input := "echo hello$USER!"
	expected := []string{"echo", "hello$USER!"}

	actual := cmdParse(input)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestCmdParseQuotedArgs(t *testing.T) {
	input := `echo "Hello World"`
	expected := []string{"echo", `"Hello`, `World"`}

	actual := cmdParse(input)

	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func trimWhite(input string) string {
	return strings.TrimSpace(input)
}

func TestTrimWhiteNoWhitespace(t *testing.T) {
	input := "ls -a"
	expected := "ls -a"

	actual := trimWhite(input)

	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}

func TestTrimWhiteLeadingWhitespace(t *testing.T) {
	input := "  ls -a"
	expected := "ls -a"

	actual := trimWhite(input)

	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}

func TestTrimWhiteTrailingWhitespace(t *testing.T) {
	input := "ls -a  "
	expected := "ls -a"

	actual := trimWhite(input)

	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}

func TestChangeDir(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get home directory: %v", err)
	}

	_, err = changeDir(homeDir)
	if err != nil {
		t.Errorf("Failed to change to home directory: %v", err)
	}

	cwd, _ := os.Getwd()
	if cwd != homeDir {
		t.Errorf("Expected directory %s, got %s", homeDir, cwd)
	}
}

func TestExecuteCommandLs(t *testing.T) {
	cmd := exec.Command("ls")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to run ls command: %v", err)
	}

	output := out.String()
	if len(output) == 0 {
		t.Errorf("Expected ls output, but got empty string")
	}
}

func TestExecuteInvalidCommand(t *testing.T) {
	cmd := exec.Command("nonexistentcommand")

	err := cmd.Run()
	if err == nil {
		t.Errorf("Expected error for non-existent command, but got none")
	}
}
