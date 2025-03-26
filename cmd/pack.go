package cmd

import (
	"github.com/spf13/cobra"
)

// packCmd represents the pack command
var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack your package",
	Run: func(cmd *cobra.Command, args []string) {
		info("Pack is not ready yet. Come back soon!")
	},
}

func init() {
	rootCmd.AddCommand(packCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// packCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// packCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
