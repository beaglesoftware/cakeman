package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

const clientID = "Ov23li1y2zrP9nlfgT0b"

// Helper function to convert string to int safely
func atoi(str string) int {
	val, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return val
}

func token_filepath() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		printerror("Failed to get home directory")
	}

	path := homedir + "/.cman/token.json"
	_, err = os.Stat(homedir + "/.cman/")
	if os.IsNotExist(err) {
		os.MkdirAll(homedir+"/.cman/", 0755)
	}

	return path
}

const deviceCodeURL = "https://github.com/login/device/code"
const tokenURL = "https://github.com/login/oauth/access_token"

// Struct for the Device Flow response
type DeviceCodeResponse struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
	Interval        int    `json:"interval"`
	ExpiresIn       int    `json:"expires_in"`
}

// Struct for the access token response
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func requestDeviceCode() (*DeviceCodeResponse, error) {
	// Define the request body
	data := map[string]string{
		"client_id": clientID,
		"scope":     "repo user", // Adjust the scopes as needed
	}
	jsonData, _ := json.Marshal(data)

	// Create the HTTP request
	resp, err := http.Post(deviceCodeURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	// fmt.Println("Raw Response:", string(body))

	// Check if GitHub returned an error
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "GitHub API error: %d %s\n", resp.StatusCode, body)
		fmt.Println("Raw Response:", string(body)) // 👀 Print raw response
		os.Exit(1)
	}

	// Parse the URL-encoded response body
	values, err := url.ParseQuery(string(body))
	if err != nil {
		fmt.Println("Error parsing URL-encoded response:", err)
		os.Exit(1)
	}

	// Extract the values from the parsed query string
	deviceResp := DeviceCodeResponse{
		DeviceCode:      values.Get("device_code"),
		UserCode:        values.Get("user_code"),
		VerificationURI: values.Get("verification_uri"),
		Interval:        atoi(values.Get("interval")),
		ExpiresIn:       atoi(values.Get("expires_in")),
	}

	// Print the user verification details
	// fmt.Println("User Code:", result["user_code"])
	// fmt.Println("Verification URL:", result["verification_uri"])
	// fmt.Println("Device Code:", result["device_code"])
	// fmt.Println("Expires In:", result["expires_in"], "seconds")
	// fmt.Println("Poll Interval:", result["interval"], "seconds")

	return &deviceResp, nil
}

// Poll GitHub for an access token
func pollForToken(deviceCode string, interval int) (*TokenResponse, error) {
	for {
		data := fmt.Sprintf("client_id=%s&device_code=%s&grant_type=urn:ietf:params:oauth:grant-type:device_code", clientID, deviceCode)
		req, err := http.NewRequest("POST", tokenURL, bytes.NewBufferString(data))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var tokenResp TokenResponse
		body, _ := io.ReadAll(resp.Body)

		if err := json.Unmarshal(body, &tokenResp); err == nil && tokenResp.AccessToken != "" {
			return &tokenResp, nil
		}

		// If authentication is still pending, wait and retry
		fmt.Println("Waiting for user to authenticate...")
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

// Save the token securely
func saveToken(token *TokenResponse) error {
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return os.WriteFile(token_filepath(), data, 0600)
}

// Load token from file
func loadToken() (*TokenResponse, error) {
	data, err := os.ReadFile(token_filepath())
	if err != nil {
		return nil, err
	}

	var token TokenResponse
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, err
	}

	return &token, nil
}
