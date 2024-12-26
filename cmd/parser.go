package cmd

import (
	"errors"
	"fmt"
	"strings"
)

// Command represents a registered command
type Command struct {
	Name    string
	Handler func(args []string) error
}

// CommandRegistry holds all registered commands
var CommandRegistry = map[string]Command{}

// RegisterCommand adds a command to the registry
func RegisterCommand(name string, handler func(args []string) error) {
	CommandRegistry[name] = Command{Name: name, Handler: handler}
}

// ParseAndExecute parses the input and executes the appropriate command
func ParseAndExecute(input string) (string, error) {
	// Tokenize input
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return "", errors.New("no command provided")
	}

	// Extract command and arguments
	commandName := tokens[0]
	args := tokens[1:]

	// Lookup the command
	command, exists := CommandRegistry[commandName]
	if !exists {
		return "", fmt.Errorf("unknown command: %s", commandName)
	}

	// Execute the command handler
	return command.Handler(args), nil
}
