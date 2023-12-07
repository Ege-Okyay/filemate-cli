package helpers

func ExtractArguments(args []string) []interface{} {
	if len(args) <= 1 {
		return nil
	}

	var extractedArgs []interface{}
	for _, arg := range args[1:] {
		extractedArgs = append(extractedArgs, arg)
	}

	return extractedArgs
}
