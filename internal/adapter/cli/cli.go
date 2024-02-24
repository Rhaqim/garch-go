package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Rhaqim/garch-go/internal/app/domain"
)

// CLIInterface provides methods for interacting with the command-line
type CLIInterface interface {
	// Start starts the CLI
	Start(config *domain.ProjectConfig)
	//Usage prints the usage message
	Usage()
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

// CLI represents the CLI adapter
type CLI struct{}

// NewCLI creates a new instance of CLI
func NewCLI() CLIInterface {
	return &CLI{}
}

func (c *CLI) Start(config *domain.ProjectConfig) {

	genCMD := flag.NewFlagSet("gen", flag.ExitOnError)

	genCMD.StringVar(&config.Title, "title", "", "Title of the project")
	genCMD.StringVar(&config.Author, "author", "", "Author of the project")
	genCMD.StringVar(&config.DbType, "db", "", "Database type")
	genCMD.StringVar(&config.Arch, "arch", "", "Architecture type")

	flag.Parse()

	genCMD.Parse(os.Args[2:])

	if config.Arch == "" {
		config.Arch = c.PromptOptions("Architecture", []string{"clean", "onion"})
	}

	if config.Title == "" {
		config.Title = c.Prompt("Title")
	}

	if config.Author == "" {
		config.Author = c.Prompt("Author")
	}

	if config.DbType == "" {
		config.DbType = c.PromptOptions("Database type", []string{"sqlite", "mysql", "postgres"})
	}

}

func (c *CLI) Usage() {
	c.Display("Usage: garch-go [command] [arguments]")
}

func (c *CLI) InvalidArgs() {
	if len(os.Args) < 2 {
		c.Display("Usage: garch-go [command] [arguments]")
		return
	}

	c.Display("Unknown command:", os.Args[1])
}

func (c *CLI) Prompt(prompt string) string {
	var input string
	fmt.Printf("%s: ", prompt)
	fmt.Scanln(&input)
	return input
}

func (c *CLI) PromptOptions(prompt string, options []string) string {
	fmt.Println(prompt)
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	var input int
	fmt.Scanln(&input)
	return options[input-1]
}

func (c *CLI) Bool(prompt string) bool {
	var input string
	fmt.Printf("%s [y/n]: ", prompt)
	fmt.Scanln(&input)

	if input != "y" && input != "n" {
		fmt.Println("Invalid input. Please enter y or n")
		return c.Bool(prompt)
	}

	return input == "y"

}

func (c *CLI) Display(a ...interface{}) {
	fmt.Println(a...)
}
