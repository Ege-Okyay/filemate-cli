package main

import (
	"fmt"
	"log"

	"github.com/Ege-Okyay/filemate-cli/cli"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("ERROR WHILE LOADING ENV VARIABLES: %s", err)
		return
	}

	err = cli.CliProgram()
	if err != nil {
		fmt.Printf("CLI ERROR: %s", err)
		return
	}
}
