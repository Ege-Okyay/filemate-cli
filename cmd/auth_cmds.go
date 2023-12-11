package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Ege-Okyay/filemate-cli/helpers"
)

func SignUpCommand(args ...interface{}) {
	username := args[0].([]interface{})[0].(string)
	email := args[1].([]interface{})[0].(string)
	password := args[2].([]interface{})[0].(string)

	values := map[string]string{
		"username": username,
		"email":    email,
		"password": password,
	}

	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	res, err := helpers.SendHttpRequest("/auth/sign-up", "POST", json_data)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var resFormat map[string]interface{}

	json.NewDecoder(res.Body).Decode(&resFormat)

	var success bool
	if resFormat["error"] != nil {
		success = false
	} else {
		success = true
	}

	if success {
		fmt.Println(resFormat["message"])
	} else {
		fmt.Println(resFormat["error"])
	}
}

func LoginCommand(args ...interface{}) {
	identifier := args[0].([]interface{})[0].(string)
	password := args[1].([]interface{})[0].(string)

	values := map[string]string{
		"identifier": identifier,
		"password":   password,
	}

	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	res, err := helpers.SendHttpRequest("/auth/login", "POST", json_data)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var resFormat map[string]interface{}

	json.NewDecoder(res.Body).Decode(&resFormat)

	var success bool
	if resFormat["error"] != nil {
		success = false
	} else {
		success = true
	}

	if success {
		fmt.Println(resFormat["message"])
		fmt.Println("---DEBUG--- TOKEN : ", resFormat["token"])
	} else {
		fmt.Println(resFormat["error"])
	}
}
