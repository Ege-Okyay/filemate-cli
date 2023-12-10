package helpers

import (
	"fmt"
	"reflect"

	"github.com/Ege-Okyay/filemate-cli/structs"
)

var commandMap = []structs.Command{}

func RegisterCommand(name string, fn interface{}, desc string, usage string, requiresArg bool) {
	newCommand := structs.Command{
		Name:        name,
		Fn:          fn,
		Desc:        desc,
		Usage:       usage,
		RequiresArg: requiresArg,
	}
	commandMap = append(commandMap, newCommand)
}

func GetAllCommands() []structs.Command {
	return commandMap
}

func FindCommandByName(name string) (*structs.Command, error) {
	for _, cmd := range commandMap {
		if cmd.Name == name {
			return &cmd, nil
		}
	}

	return nil, fmt.Errorf("Command not found: %s", name)
}

func CallCommand(cmd *structs.Command, args ...interface{}) {
	value := reflect.ValueOf(cmd.Fn)

	if value.Kind() != reflect.Func {
		panic(fmt.Errorf("Not a valid function"))
	}

	var inputValues []reflect.Value

	if len(args) > 0 {
		for _, arg := range args {
			inputValues = append(inputValues, reflect.ValueOf(arg))
		}
	}

	if cmd.RequiresArg && len(inputValues) == 0 {
		fmt.Println("Correct usage:", cmd.Usage)
		return
	}

	value.Call(inputValues)
}
