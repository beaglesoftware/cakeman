package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

func DownloadFile(url, filePath string) error {
	// Get the data from the URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Copy the response body to the file
	_, err = io.Copy(out, resp.Body)
	return err
}

func GetURLContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compiles dependencies and app itself",
	Run: func(cmd *cobra.Command, args []string) {
		info("Building started!")
		os.RemoveAll("headers")
		file, err := os.ReadFile("Cake.cman")
		if err != nil {
			printerror("Error reading file: " + err.Error())
			os.Exit(2)
		}

		var config map[string]map[string]interface{}

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

		dependencies := config["Dependencies"]
		for depName, depVersion := range dependencies {
			// Currently, Cakeman only supports the latest version of a cake
			depVersion = "latest"
			info(fmt.Sprintf("Downloading %s version %s", depName, depVersion))
			os.MkdirAll("headers/"+depName, 0755)
			var firstchar string
			for _, char := range depName {
				firstchar = string(char)
				// fmt.Println(firstchar)
				break
			}
			cakefileContent, err := GetURLContent("https://raw.githubusercontent.com/beaglesoftware/cakes/refs/heads/main/manifests/" + firstchar + "/" + depName + ".cman")
			// fmt.Println(cakefileContent)
			if err != nil {
				printerror("An error occurred: " + err.Error())
			}
			var depConfig struct {
				Package struct {
					Repository string `json:"Repository"`
					// Main       string `json:"Main"`
				} `json:"Package"`
			}
			if cakefileContent == "" {
				printerror("Received empty content for " + depName)
				continue
			}
			err = json.Unmarshal([]byte(cakefileContent), &depConfig)
			if err != nil {
				// fmt.Println(depConfig)
				printerror("Failed to unmarshal dependency config: " + err.Error())
				continue
			}
			repository := depConfig.Package.Repository
			// fmt.Println(repository)
			outFile := "headers/" + depName + "/"
			repo_url := "https://github.com/" + repository + ".git"
			run_command("git", "clone", repo_url, outFile)
			// fmt.Println(url)
			if err != nil {
				printerror("An error occurred: " + err.Error())
			}
			success("Dependency " + depName + " installed successfully!")
		}

		cakeName := config["Package"]["Name"].(string)
		os.MkdirAll("./cman-build/", 0755)

		var app string
		if mainFile, ok := config["Package"]["Main"].(string); ok && strings.HasSuffix(mainFile, ".c") {
			if runtime.GOOS == "darwin" {
				app = "/usr/bin/clang"
			} else if runtime.GOOS == "linux" {
				app = "gcc"
			} else if runtime.GOOS == "windows" {
				app = "cl"
			}
		} else if mainFile, ok := config["Package"]["Main"].(string); ok && strings.HasSuffix(mainFile, ".cpp") {
			if runtime.GOOS == "darwin" {
				app = "/usr/bin/clang++"
			} else if runtime.GOOS == "linux" {
				app = "g++"
			} else if runtime.GOOS == "windows" {
				app = "cl"
			}
		}

		_, err = os.Stat(config["Package"]["Main"].(string))
		if os.IsNotExist(err) {
			printerror(config["Package"]["Main"].(string) + " does not exist!")
		} else if err != nil {
			fmt.Println("Error accessing file:", err)
		}

		arg0 := config["Package"]["Main"].(string)
		var arg1 string
		currentdir, err := os.Getwd()
		if err != nil {
			printerror("Failed to get current directory: " + err.Error())
		}
		if runtime.GOOS == "darwin" {
			arg1 = "-o"
		} else if runtime.GOOS == "linux" {
			arg1 = "-o"
		} else if runtime.GOOS == "windows" {
			arg1 = "/Fe" + currentdir + "/cman-build/" + cakeName + ".exe"
		}
		var arg2 string
		if runtime.GOOS == "darwin" {
			arg2 = currentdir + "/cman-build/" + cakeName
		} else if runtime.GOOS == "linux" {
			arg2 = currentdir + "/cman-build/" + cakeName
		} else if runtime.GOOS == "windows" {
			arg2 = "/Fo" + currentdir + "/cman-build/" + cakeName + ".obj"
		}
		run_command(app, arg0, arg1, arg2)

		success(cakeName + " installed successfully!")
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
