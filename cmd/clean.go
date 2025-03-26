package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up by deleting build files",
	Long:  `Remove build files and headers. Note that they will be created when you compile the app again`,
	Run: func(cmd *cobra.Command, args []string) {
		err := os.RemoveAll("headers")
		if err != nil {
			printerror("Failed to remove folder 'headers': " + err.Error())
		}
		err = os.RemoveAll("cman-build")
		if err != nil {
			printerror("Failed to remove folder 'cman-build': " + err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
