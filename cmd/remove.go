package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a package from Cake.cman",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.ReadFile("Cake.cman")
		if err != nil {
			printerror("Error reading file: " + err.Error())
			os.Exit(2)
		}

		// Check if arguments are provided
		if len(args) == 0 {
			printerror("Not all required arguments provided!")
			os.Exit(4)
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

		if details, ok := config["details"].(map[string]interface{}); ok {
			delete(details, "active")
		} else {
			printerror(args[0] + " does not exists")
		}

		delete(config, args[0])
		fmt.Println(args[0])

		modifiedJSON, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		file2, err := os.Create("Cake.cman")

		if err != nil {
			printerror("Failed to create file Cake.cman")
		}
		file2.Write(modifiedJSON)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
