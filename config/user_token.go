package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// userConfigFile is the file path for storing user configuration.
const userConfigFile = "config/user_config.txt"

// userToken represents the user token used for authentication.
var userToken string

// init is called during package initialization to read the user token from the configuration file.
func init() {
	readUserTokenFromFile()
}

// GetUserToken retrieves the current user token.
func GetUserToken() string {
	return userToken
}

// SetUserToken sets a new user token and writes it to the configuration file.
func SetUserToken(newToken string) {
	userToken = newToken
	writeUserTokenToFile()
}

// readUserTokenFromFile reads the user token from the configuration file.
func readUserTokenFromFile() {
	content, err := os.ReadFile(userConfigFile)
	if err != nil {
		log.Fatal("Error reading user config file: ", err)
	}

	// Split the content by "=" to extract the user token
	parts := strings.Split(string(content), "=")
	if len(parts) != 2 {
		log.Fatal("Unexpected config format")
	}

	userToken = parts[1]
}

// writeUserTokenToFile writes the user token to the configuration file.
func writeUserTokenToFile() {
	// Create data string in the format "USER_TOKEN=<token>"
	data := []byte(fmt.Sprintf("USER_TOKEN=%s", userToken))

	// Write the data to the user configuration file
	err := os.WriteFile(userConfigFile, data, 0644)
	if err != nil {
		log.Fatal("Error writing user token to file: ", err)
	}
}
