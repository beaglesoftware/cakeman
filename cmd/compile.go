package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compiles dependencies and app itself",
	Run: func(cmd *cobra.Command, args []string) {
		info("Building started!")
		file, err := os.ReadFile("Cake.cman")
		if err != nil {
			printerror("Error reading file: " + err.Error())
			os.Exit(2)
		}

		var config map[string]interface{}

		err = json.Unmarshal(file, &config)
		if err != nil {
			if runtime.GOOS == "windows" {
				bgred := color.New(color.BgRed).SprintFunc()
				fmt.Println(color.Output, bgred("ERROR"), "Failed to read JSON", err)
			} else {
				bgred := color.New(color.BgRed).SprintFunc()
				fmt.Println(bgred("ERROR"), "Failed to read JSON", err)
			}
		}

		for
	},
}

func init() {
	rootCmd.AddCommand(compileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// compileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// compileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
