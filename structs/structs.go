package structs

type Command struct {
	Name        string
	Fn          interface{}
	Desc        string
	Usage       string
	RequiresArg bool
}
