package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var authenticateCmd = &cobra.Command{
	Use:   "authenticate",
	Short: "Login to your GitHub account",
	Long:  `This command logins to your GitHub account for publishing your cake and...`,
	Run: func(cmd *cobra.Command, args []string) {
		info("Authenticating to GitHub...")
		// Check if we already have a token
		token, err := loadToken()
		if err == nil {
			info("Already authenticated. Token: " + token.AccessToken)
			os.Exit(0)
		}

		// Step 1: Get a device code
		deviceCodeResp, err := requestDeviceCode()
		if err != nil {
			printerror("Error requesting device code:" + err.Error())
		}
		info(fmt.Sprintf("Open %s and enter the code: %s\n", deviceCodeResp.VerificationURI, deviceCodeResp.UserCode))

		// Step 2: Poll for authentication
		token, err = pollForToken(deviceCodeResp.DeviceCode, deviceCodeResp.Interval)
		if err != nil {
			printerror("Error getting token:" + err.Error())
		}

		// Step 3: Save token securely
		if err := saveToken(token); err != nil {
			log.Fatal("Error saving token:", err)
		}

		success("Authentication successful! Token saved.")
	},
}

func init() {
	rootCmd.AddCommand(authenticateCmd)
}
