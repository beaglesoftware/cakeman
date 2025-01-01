package cmd

import (
	"fmt"
	"os"
	"runtime"

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

func hint(msg string) {
	if runtime.GOOS == "windows" {
		magenta := color.New(color.BgMagenta).SprintFunc()
		fmt.Println(color.Output, magenta("HINT"), msg)
	} else {
		magenta := color.New(color.BgMagenta).SprintFunc()
		fmt.Println(magenta("HINT"), msg)
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

type Config struct {
	Package      Package                `toml:"package"`
	Dependencies map[string]interface{} `toml:"dependencies"`
}

type Package struct {
	Name        string `toml:"name"`
	Description string `toml:"description"`
	License     string `toml:"license"`
	Main        string `toml:"main"`
	Author      string `toml:"author"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cakeman",
	Short: "Knock knock! Your C/C++ package delivered!",
	Long:  `The missing package manager for C and C++`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
