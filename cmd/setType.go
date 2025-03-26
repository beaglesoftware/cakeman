/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// setTypeCmd represents the setType command
var setTypeCmd = &cobra.Command{
	Use:   "set-type",
	Short: "Set the type of the cake",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		var caketype string
		supportedType := []string{"bin", "lib"}
		if len(args) < 2 {
			printerror("You must specify a cake type and your cake's name.")
			printerror("Usage: cman init [CAKENAME] [CAKETYPE]")
			os.Exit(1)
		}
		cakename := args[0]
		caketype = args[1]
		found := false
		for _, t := range supportedType {
			if t == caketype {
				found = true
				break
			}
		}
		if !found {
			cmd.Println("Unsupported cake type. Supported types are: bin, lib")
			os.Exit(1)
		}

		if caketype == "bin" {
			err := os.Rename(cakename + ".cman", "Cake.cman")
			if err != nil {
				printerror("An error occurred when trying to rename file: " + err.Error())
			}
		} else if caketype == "lib" {
			err := os.Rename("Cake.cman", cakename + ".cman")
			if err != nil {
				printerror("An error occurred when trying to rename file: " + err.Error())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setTypeCmd)
}
