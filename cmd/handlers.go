package cmd

import (
	"fmt"
)

// HandleCake processes the "cake" command
func HandleCake(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("cake command requires exactly 1 argument: <name>")
	}
	fmt.Printf("Baking a cake named: %s\n", args[0])
	return nil
}

// HandleBuild processes the "build" command
func HandleBuild(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("build command requires exactly 1 argument: <filename>")
	}
	fmt.Printf("Building file: %s\n", args[0])
	return nil
}

// HandleModule processes the "module" command
func HandleModule(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("module command requires exactly 1 argument: <module>")
	}
	fmt.Printf("Loading module: %s\n", args[0])
	return nil
}
