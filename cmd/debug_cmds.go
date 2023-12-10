package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Ege-Okyay/filemate-cli/helpers"
)

func HelloCmd(args []interface{}) {
	resMsg := fmt.Sprintf("Hello, %v", args[0])

	fmt.Println(resMsg)
}

func HealthCheckCmd(args []interface{}) {
	res, err := helpers.SendHttpRequest("/health-check", "GET", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println("Status: ", res.Status)
}

func PostCheckCmd(args []interface{}) {
	values := map[string]string{"param": args[0].(string)}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	res, err := helpers.SendHttpRequest("/post-check", "POST", json_data)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var resp map[string]interface{}

	json.NewDecoder(res.Body).Decode(&resp)

	fmt.Println(resp)
}
