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

// File represents a file entity with metadata.
type File struct {
	ID       string `json:"id"`       // ID is the unique identifier for the file.
	UserID   string `json:"userId"`   // UserID is the identifier of the user who uploaded the file.
	FileName string `json:"fileName"` // FileName is the name of the file.
	FilePath string `json:"filePath"` // FilePath is the path where the file is stored.
	FileSize int64  `json:"fileSize"` // FileSize is the size of the file in bytes.
	Public   bool   `json:"public"`   // Public indicates whether the file is public or private.
}
