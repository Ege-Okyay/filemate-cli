package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const userConfigFile = "config/user_config.txt"

var userToken string

func init() {
	readUserTokenFromFile()
}

func GetUserToken() string {
	return userToken
}

func SetUserToken(newToken string) {
	userToken = newToken
	writeUserTokenToFile()
}

func readUserTokenFromFile() {
	content, err := os.ReadFile(userConfigFile)
	if err != nil {
		log.Fatal("Error reading user config file: ", err)
	}

	parts := strings.Split(string(content), "=")
	if len(parts) != 2 {
		log.Fatal("Unexpected config format")
	}

	userToken = parts[1]
}

func writeUserTokenToFile() {
	data := []byte(fmt.Sprintf("USER_TOKEN=%s", userToken))
	err := os.WriteFile(userConfigFile, data, 0644)
	if err != nil {
		log.Fatal("Error writing user token to file: ", err)
	}
}
