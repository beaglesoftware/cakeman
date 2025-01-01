package cmd

import (
	"os"

	"encoding/json"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func initproj(name string, filename string, author string) {
	config := Config{
		Package: Package{
			Name:        name,
			Description: "",
			License:     "",
			Main:        filename,
			Author:      author,
		},
		Dependencies: make(map[string]interface{}),
	}
	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		printerror("Failed to convert structure to TOML")
	}
	file, err := os.Create("Cake.cman")
	if err != nil {
		printerror("Failed to save to file")
	}

	file.Write(data)
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project and write to 'Packages.cman'",
	Long:  ``,
	Args:  cobra.MaximumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			printerror("Not all required arguments provided! Pass `--help` for usage")
			os.Exit(4)
		}
		name := args[0]
		var filename string
		if len(args) > 1 {
			filename = args[1]
		} else {
			filename = "main.c"
		}

		var author string

		if len(args) > 1 {
			author = args[1]
		} else {
			author = "" // Set a default value if not provided
		}

		initproj(name, filename, author)
		color.Green("Project initialized successfully!")
		info("Read Cakefile best practice at https://github.com/beaglesoftware/cakeman/blob/main/BESTPRACTICE.md")
		info("Run 'cman build' to build the project")
		info("Run 'cman run' to run the project")
		info("Run 'cman add' to add some dependencies")
		info("Run 'cman pack' to pack the cake")
		info("Run 'cman publish' to publish the cake")
		color.HiGreen("Happy coding!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
