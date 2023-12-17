package helpers

// ExtractArguments takes a slice of command-line arguments and returns a slice of interfaces.
// It extracts arguments excluding the program name (args[0]) and returns them as a slice of interfaces.
func ExtractArguments(args []string) []interface{} {
	// Check if there are more than one argument (program name is at index 0)
	if len(args) <= 1 {
		return nil
	}

	var extractedArgs []interface{}
	// Iterate through the arguments starting from index 1 (excluding the program name)
	for _, arg := range args[1:] {
		extractedArgs = append(extractedArgs, arg)
	}

	return extractedArgs
}
