package cmd

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func printerror(msg string) {
	if runtime.GOOS == "windows" {
		red := color.New(color.BgRed).SprintFunc()
		fmt.Println(color.Output, red("ERROR"), msg)
	} else {
		red := color.New(color.BgRed).SprintFunc()
		fmt.Println(red("ERROR"), msg)
	}
}

func info(msg string) {
	if runtime.GOOS == "windows" {
		bgblue := color.New(color.BgBlue).SprintFunc()
		fmt.Println(color.Output, bgblue("INFO"), msg)
	} else {
		bgblue := color.New(color.BgBlue).SprintFunc()
		fmt.Println(bgblue("INFO"), msg)
	}
}

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
		printerror("Error writing to file: " + err.Error())
		os.Exit(11)
	}
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project and write to 'Packages.cman'",
	Long:  ``,
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			printerror("Not all required arguments provided!")
			os.Exit(4)
		}
		name := args[0]
		var filename string
		if len(args) > 1 {
			filename = args[1]
		} else {
			filename = "main.c"
		}
		initproj(name, filename)
		color.Green("Project initialized successfully!")
		info("Read Cakefile best practice at https://github.com/beaglesoftware/cakeman/blob/main/BESTPRACTICE.md")
		info("Run 'cake build' to build the project")
		info("Run 'cake run' to run the project")
		info("Run 'cake add' to add some dependencies")
		info("Run 'cake pack' to pack the cake")
		info("Run 'cake publish' to publish the cake")
		color.HiGreen("Happy coding!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().String("author", "", "Author of the project")
}
