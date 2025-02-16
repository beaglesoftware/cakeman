package cmd

import (
	"fmt"
	"os"

	"encoding/json"

	"github.com/spf13/cobra"
)

func initproj(name string, filename string, author string) {
	config := Config{
		Package: Package{
			Name:        name,
			Description: "",
			License:     "",
			Repository:  "",
			Main:        filename,
			Author:      author,
		},
		Dependencies: make(map[string]interface{}),
	}
	data, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		printerror("Failed to convert structure to TOML")
	}
	var file *os.File
	file, err = os.Create("Cake.cman")
	if err != nil {
		printerror("Failed to save to file")
	}

	file.Write(data)
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project and write to 'Cake.cman' or '{name}.cman' (if it is a library)",
	Long:  ``,
	Args:  cobra.MaximumNArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			printerror("Not all required arguments provided!\nUsage: cman init [NAME] [MAINFILENAME] [AUTHOR | optional] [--header | optional, use it if you want to initialize a library]")
			os.Exit(4)
		}
		// fmt.Println(args)
		name := args[0]
		var filename string
		if len(args) > 1 {
			filename = args[1]
		} else {
			filename = "main.c"
		}

		var author string

		if len(args) > 2 {
			author = args[2]
		} else {
			author = ""
		}

		initproj(name, filename, author)
		success("Project initialized successfully!")
		info("Read Cakefile best practice at https://github.com/beaglesoftware/cakeman/blob/main/BESTPRACTICE.md")
		info("Run 'cman build' to build the project")
		info("Run 'cman run' to run the project")
		info("Run 'cman add' to add some dependencies")
		info("Run 'cman pack' to pack the cake")
		info("Run 'cman publish' to publish the cake")
		info("Run 'cman set-type lib' to turn this cake into a library that includes some headers")
		fmt.Println("Happy coding!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
