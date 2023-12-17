package helpers

import (
	"fmt"
	"reflect"

	"github.com/Ege-Okyay/filemate-cli/structs"
)

// commandMap is a slice of Command structs to store registered commands.
var commandMap = []structs.Command{}

// RegisterCommand registers a new command with the specified properties.
// It adds the command to the command map.
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

// GetAllCommands returns a slice of all registered commands.
func GetAllCommands() []structs.Command {
	return commandMap
}

// FindCommandByName finds a command by its name in the command map.
// It returns a pointer to the Command if found, otherwise returns an error.
func FindCommandByName(name string) (*structs.Command, error) {
	for _, cmd := range commandMap {
		if cmd.Name == name {
			return &cmd, nil
		}
	}

	return nil, fmt.Errorf("Command not found: %s", name)
}

// CallCommand calls the function associated with the specified command and passes the arguments.
// It validates the number of arguments based on the command's requirements.
func CallCommand(cmd *structs.Command, args []interface{}) {
	value := reflect.ValueOf(cmd.Fn)

	// Ensure that the value is a function
	if value.Kind() != reflect.Func {
		panic(fmt.Errorf("Not a valid function"))
	}

	var inputValues []reflect.Value

	// Convert arguments to reflect values
	if len(args) > 0 {
		for _, arg := range args {
			argValue := reflect.ValueOf(arg)

			// If the argument is a string, convert it to a slice of interfaces
			if argValue.Kind() == reflect.String {
				argValue = reflect.ValueOf([]interface{}{arg})
			}

			// If the argument is a pointer, dereference it
			if argValue.Kind() == reflect.Ptr {
				argValue = argValue.Elem()
			}

			inputValues = append(inputValues, argValue)
		}
	}

	// Validate the number of arguments based on the command's requirements
	if !cmd.ArgsOptional {
		if len(inputValues) != cmd.NumberOfArgsRequired {
			fmt.Println("Incorrect number of arguments.\nCorrect usage: ", cmd.Usage)
			return
		}
	}

	// Call the function with the provided arguments
	value.Call(inputValues)
}
