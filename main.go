package main

import (
	"fmt"
	"log"

	"github.com/Ege-Okyay/filemate-cli/cli"
	"github.com/joho/godotenv"
)

// main is the entry point of the Filemate CLI application.
// It loads environment variables, initializes the CLI program, and handles errors.
func main() {
	var err error

	// Load environment variables from the .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env variables: %s", err)
		return
	}

	// Execute the CLI program
	err = cli.CliProgram()
	if err != nil {
		fmt.Printf("CLI Error : %s", err)
		return
	}
}
