package helpers

import (
	"fmt"
	"reflect"

	"github.com/Ege-Okyay/filemate-cli/structs"
)

var commandMap = []structs.Command{}

func RegisterCommand(name string, fn interface{}, desc string, usage string, category string, numberOfArgsRequired int, argsOptional bool) {
	newCommand := structs.Command{
		Name:                 name,
		Fn:                   fn,
		Desc:                 desc,
		Usage:                usage,
		NumberOfArgsRequired: numberOfArgsRequired,
		ArgsOptional:         argsOptional,
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

func CallCommand(cmd *structs.Command, args []interface{}) {
	value := reflect.ValueOf(cmd.Fn)

	if value.Kind() != reflect.Func {
		panic(fmt.Errorf("Not a valid function"))
	}

	var inputValues []reflect.Value

	if len(args) > 0 {
		for _, arg := range args {
			argValue := reflect.ValueOf(arg)

			if argValue.Kind() == reflect.String {
				argValue = reflect.ValueOf([]interface{}{arg})
			}

			if argValue.Kind() == reflect.Ptr {
				argValue = argValue.Elem()
			}

			inputValues = append(inputValues, argValue)
		}
	}

	// Check if the inputValues are less than the required amount
	if !cmd.ArgsOptional {
		if len(inputValues) < cmd.NumberOfArgsRequired {
			fmt.Println("Number of arguments are less than required amount.\nCorrect usage: ", cmd.Usage)
			return
		} else if len(inputValues) > cmd.NumberOfArgsRequired {
			fmt.Println("Number of arguments are more than required amount.\nCorrect usage: ", cmd.Usage)
			return
		}
	}

	value.Call(inputValues)
}
