package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add package",
	Long:  ``,
	Args:  cobra.MaximumNArgs(1),
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

		var firstchar string

		for _, char := range args[0] {
			firstchar = string(char)
			break
		}

		info("Getting cakes...")

		url := "https://github.com/beaglesoftware/cakes/blob/main/manifests/" + firstchar + "/" + args[0] + ".cman"
		response, err := http.Get(url)
		if err != nil {
			printerror("Failed to send request to beaglesoftware/cakes GitHub repo:" + err.Error())
			os.Exit(1)
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			printerror("Got a non-200 HTTP status code: " + string(rune(response.StatusCode)))
			if response.StatusCode == http.StatusNotFound {
				hint("Error code is 404. Usually 404 error code will be returned if cake doesn't exists. Try adding an existing one")
			}
			os.Exit(15)
		}

		var config Config

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

		config.Dependencies[args[0]] = "latest"

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
	rootCmd.AddCommand(addCmd)
}
