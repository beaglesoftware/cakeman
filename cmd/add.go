package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add package",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		buf, err := os.ReadFile("Cake.cman")
		if err != nil {
			printerror("Error reading file: " + err.Error())
			os.Exit(2)
		}
		inputs := strings.Split(string(buf), "\n") // Split file content into lines

		// Check if arguments are provided
		if len(args) == 0 {
			printerror("Not all required arguments provided!")
			os.Exit(4)
		}

		// Loop through the inputs and process each
		for _, input := range inputs {
			if input != "" { // Ensure we're not processing empty lines
				err = ParseAndExecute(input)
				if err != nil {
					printerror(err.Error())
					os.Exit(5)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
