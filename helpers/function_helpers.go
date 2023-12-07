package helpers

import (
	"fmt"
	"reflect"
)

var functionMap = make(map[string]interface{})

func RegisterCommand(name string, fn interface{}) {
	functionMap[name] = fn
}

func FindCommandByName(name string) (interface{}, error) {
	fn, found := functionMap[name]
	if !found {
		return nil, fmt.Errorf("Function not found: %s", name)
	}
	return fn, nil
}

func CallCommand(fn interface{}, args ...any) error {
	value := reflect.ValueOf(fn)

	if value.Kind() != reflect.Func {
		return fmt.Errorf("Not a valid function")
	}

	var inputValues []reflect.Value
	for _, arg := range args {
		inputValues = append(inputValues, reflect.ValueOf(arg))
	}

	value.Call(inputValues)

	return nil
}
