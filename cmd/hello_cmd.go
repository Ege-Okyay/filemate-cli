package cmd

import (
	"fmt"
)

func HelloCmd(args []interface{}) {
	if args[0] == "-h" || args[0] == "--help" {
		fmt.Println("Usage : filemate hello [name]")
		return
	}

	resMsg := fmt.Sprintf("Hello, %v", args[0])

	fmt.Println(resMsg)
}
