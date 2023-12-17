package structs

// Command represents the structure of a command in the Filemate CLI.
type Command struct {
	Name                 string      // Name of the command
	Fn                   interface{} // Function associated with the command
	Desc                 string      // Description of the command
	Usage                string      // Usage information for the command
	NumberOfArgsRequired int         // Number of arguments required by the command
	ArgsOptional         bool        // Flag indicating whether arguments are optional for the command
}
