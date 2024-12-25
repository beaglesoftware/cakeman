package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func initproj(name string, filename string) {
	file, err := os.Create("Cake.cman")
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		fmt.Println(color.Output, red("Error creating file: "), err)
		os.Exit(11)
	}

	defer file.Close()
	var builder strings.Builder
	builder.WriteString("cake " + name + "\n")
	builder.WriteString("build " + filename + "\n")
	builder.WriteString("")
	content := builder.String()
	_, err = file.WriteString(content)
	if err != nil {
		red := color.New(color.FgRed).SprintFunc()
		fmt.Println(color.Output, red("Error writing to file: "), err)
		os.Exit(11)
	}
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project and write to 'Packages.cman'",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		filename := "main.c"
		if len(args) > 0 {
			filename = args[1]
		}
		initproj(name, filename)
		color.Green("Project initialized successfully!")
		bgblue := color.New(color.BgBlue).SprintFunc()
		fmt.Println(color.Output, bgblue("INFO"), "Read Cakefile best practice at https://github.com/beaglesoftware/cakeman/blob/main/BESTPRACTICE.md")
		fmt.Println(color.Output, bgblue("INFO"), "Run 'cake build' to build the project")
		fmt.Println(color.Output, bgblue("INFO"), "Run 'cake run' to run the project")
		fmt.Println(color.Output, bgblue("INFO"), "Run 'cake add' to add some dependencies")
		color.HiGreen("Happy coding!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().String("author", "", "Author of the project")
}
