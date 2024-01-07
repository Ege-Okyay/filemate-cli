package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Ege-Okyay/filemate-cli/config"
	"github.com/Ege-Okyay/filemate-cli/helpers"
)

// SignUpCommand handles the "sign-up" command, allowing users to create an account with a given username, email, and password.
func SignUpCommand(args ...interface{}) {
	// Check if the user is already logged in
	if config.GetUserToken() != "" {
		fmt.Println("Already logged in")
		return
	}

	// Extract username, email, and password from the arguments
	username := args[0].([]interface{})[0].(string)
	email := args[1].([]interface{})[0].(string)
	password := args[2].([]interface{})[0].(string)

	// Prepare data for the HTTP request
	values := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
	}

	// Convert data to JSON
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	// Send HTTP request to sign up the user
	res, err := helpers.SendHttpRequest("/auth/sign-up", "POST", jsonData, nil, "", "")
	if err != nil {
		log.Fatal(err)
	}

	// Decode the response JSON
	var resFormat map[string]interface{}
	json.NewDecoder(res.Body).Decode(&resFormat)

	// Check if the sign-up was successful
	if resFormat["error"] == nil {
		fmt.Println(resFormat["message"])
	} else {
		fmt.Println(resFormat["error"])
	}
}

// LoginCommand handles the "login" command, allowing users to log in with a username or email and password.
func LoginCommand(args ...interface{}) {
	// Check if the user is already logged in
	if config.GetUserToken() != "" {
		fmt.Println("Already logged in")
		return
	}

	// Extract identifier (username or email) and password from the arguments
	identifier := args[0].([]interface{})[0].(string)
	password := args[1].([]interface{})[0].(string)

	// Prepare data for the HTTP request
	values := map[string]string{
		"identifier": identifier,
		"password":   password,
	}

	// Convert data to JSON
	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	// Send HTTP request to log in the user
	res, err := helpers.SendHttpRequest("/auth/login", "POST", jsonData, nil, "", "")
	if err != nil {
		log.Fatal(err)
	}

	// Decode the response JSON
	var resFormat map[string]interface{}
	json.NewDecoder(res.Body).Decode(&resFormat)

	// Check if the login was successful
	if resFormat["error"] == nil {
		fmt.Println(resFormat["message"])

		// Set user token in the configuration
		config.SetUserToken(resFormat["token"].(string))
	} else {
		fmt.Println(resFormat["error"])
	}
}

// LogoutCommand handles the "logout" command, allowing users to log out.
func LogoutCommand(args ...interface{}) {
	// Check if the user is logged in
	if config.GetUserToken() == "" {
		fmt.Println("You should log in first")
		return
	}

	// Clear user token in the configuration
	config.SetUserToken("")
	fmt.Println("Successfully logged out")
}
