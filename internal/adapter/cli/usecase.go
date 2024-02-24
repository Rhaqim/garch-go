package cli

import "github.com/Rhaqim/garch-go/internal/app/domain"

// CLIInterface provides methods for interacting with the command-line
type CLIInterface interface {
	// Start starts the CLI
	Start(config *domain.ProjectConfig)
	//Usage prints the usage message
	Usage()
	// InvalidArgs prints the invalid arguments message
	InvalidArgs()
	// HandleArgs handles the command-line arguments
	HandleArgs(config *domain.ProjectConfig)
	// Prompt asks the user for input
	Prompt(prompt string) string
	// PromptOptions asks the user to choose from a list of options
	PromptOptions(prompt string, options []string) string
	// Bool asks the user for a boolean input, Yes or No
	Bool(prompt string) bool
	// Println prints a line to the console
	Display(a ...interface{})
}
