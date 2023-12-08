package cmd

import (
	"fmt"
	"log"

	"github.com/Ege-Okyay/filemate-cli/helpers"
)

func HelloCmd(args []interface{}) {
	resMsg := fmt.Sprintf("Hello, %v", args[0])

	fmt.Println(resMsg)
}

func HealthCheckCmd(args []interface{}) {
	res, err := helpers.SendGETRequest("/health-check")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Status: ", res.Status)
}
