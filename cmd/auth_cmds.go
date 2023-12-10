package cmd

import "fmt"

func SignUpCommand(args []interface{}) {

}

func LoginCommand(args []interface{}) {
	identifier := args[0]
	password := args[1]

	fmt.Printf("---DEBUG--- IDENTIFIER : %s, PASSWORD : %s", identifier, password)
}
