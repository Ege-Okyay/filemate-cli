package cmd

import (
	"fmt"
)

func HelloCmd(args []interface{}) {
	resMsg := fmt.Sprintf("Hello, %v", args[0])

	fmt.Println(resMsg)
}
