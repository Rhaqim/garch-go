package cli

// CLIInterface provides methods for interacting with the command-line
type CLIInterface interface {
	// InvalidArgs prints the invalid arguments message
	InvalidArgs()
	// Prompt asks the user for input
	Prompt(prompt string) string
	// PromptOptions asks the user to choose from a list of options
	PromptOptions(prompt string, options []string) string
	// Bool asks the user for a boolean input, Yes or No
	Bool(prompt string) bool
	// Println prints a line to the console
	Display(a ...interface{})
}
