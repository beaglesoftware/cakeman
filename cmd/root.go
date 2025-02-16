package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

func printerror(msg string) {
	if runtime.GOOS == "windows" {
		red := color.New(color.BgRed).SprintFunc()
		fgwhite := color.New(color.FgWhite).SprintFunc()
		fmt.Println(color.Output, red(fgwhite("ERROR")), msg)
	} else {
		red := color.New(color.BgRed).SprintFunc()
		fgwhite := color.New(color.FgWhite).SprintFunc()
		fmt.Println(red(fgwhite("ERROR")), msg)
	}
}

func hint(msg string) {
	if runtime.GOOS == "windows" {
		magenta := color.New(color.BgMagenta).SprintFunc()
		fgwhite := color.New(color.FgWhite).SprintFunc()
		fmt.Println(color.Output, magenta(fgwhite("HINT")), msg)
	} else {
		magenta := color.New(color.BgMagenta).SprintFunc()
		fgwhite := color.New(color.FgWhite).SprintFunc()
		fmt.Println(magenta(fgwhite("HINT")), msg)
	}
}

func info(msg string) {
	if runtime.GOOS == "windows" {
		bgblue := color.New(color.BgBlue).SprintFunc()
		fgwhite := color.New(color.FgWhite).SprintFunc()
		fmt.Println(color.Output, bgblue(fgwhite("INFO")), msg)
	} else {
		bgblue := color.New(color.BgBlue).SprintFunc()
		fgwhite := color.New(color.FgWhite).SprintFunc()
		fmt.Println(bgblue(fgwhite("INFO")), msg)
	}
}

func success(msg string) {
	if runtime.GOOS == "windows" {
		bghigreen := color.New(color.BgHiGreen).SprintFunc()
		fgwhite := color.New(color.FgWhite).SprintFunc()
		fmt.Println(color.Output, bghigreen(fgwhite("SUCCESS")), msg)
	} else {
		bghigreen := color.New(color.BgHiGreen).SprintFunc()
		fgwhite := color.New(color.FgWhite).SprintFunc()
		fmt.Println(bghigreen(fgwhite("SUCCESS")), msg)
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
	Repository  string `toml:"repository"`
	Main        string `toml:"main"`
	Author      string `toml:"author"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cakeman",
	Short: "Knock knock! Your C/C++ package delivered!",
	Long: `Cakeman is a package manager for C/C++.
	It does anything, note that it is currently in preview. Contributions are welcome.`,
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

func run_command(app string, arg0 string, arg1 string, arg2 string) {
	var command string
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		command = app + " " + arg0 + " " + arg1 + " " + arg2
		cmd = exec.Command(command)
	} else {
		command = app + " " + arg0 + " " + arg1 + " " + arg2
		script := "cman-build/compile.sh"
		content := "#!/bin/bash\n" + command
		// Write script to a temporary file
		err := os.WriteFile(script, []byte(content), 0755)
		if err != nil {
			fmt.Println("Error writing script:", err)
			return
		}

		cmd = exec.Command("/bin/bash", "-c", "./cman-build/compile.sh")
	}
	// os.Exit(0)
	current, err := os.Getwd()
	if err != nil {
		printerror("Error when getting current directory: " + err.Error())
	}
	cmd.Dir = current
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		printerror("Failed to get stderr pipe: " + err.Error())
		os.Exit(1)
	}

	if err := cmd.Start(); err != nil {
		printerror("Failed to start command: " + err.Error())
		os.Exit(1)
	}

	stderrBytes, err := io.ReadAll(stderrPipe)
	if err != nil {
		printerror("Failed to read stderr: " + err.Error())
		os.Exit(1)
	}

	if err := cmd.Wait(); err != nil {
		printerror(app + " " + arg0 + " " + arg1 + " " + arg2 + " exited with error: " + err.Error() + "\n" + string(stderrBytes))
		os.Exit(1)
	}
}
