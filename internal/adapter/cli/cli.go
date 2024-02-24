package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Rhaqim/garch-go/internal/app/domain"
	"github.com/Rhaqim/garch-go/internal/app/usecase"
)

// CLIInterface provides methods for interacting with the command-line
type CLIInterface interface {
	//Start starts the CLI
	Start()
	// Prompt asks the user for input
	Prompt(prompt string) string
	// PromptOptions asks the user to choose from a list of options
	PromptOptions(prompt string, options []string) string
	// Println prints a line to the console
	Display(a ...interface{})
	// HandleArgs handles the command-line arguments
	HandleArgs()
	// InvalidArgs prints the invalid arguments message
	InvalidArgs()
}

// CLI represents the CLI adapter
type CLI struct {
	project usecase.ProjectUseCase
}

// NewCLI creates a new instance of CLI
func NewCLI(project usecase.ProjectUseCase) CLIInterface {
	return &CLI{
		project: project,
	}
}

func (c *CLI) Start() {
	c.Display("Welcome to Garch!")

	config := domain.ProjectConfig{}

	genCMD := flag.NewFlagSet("gen", flag.ExitOnError)

	genCMD.StringVar(&config.Title, "title", "", "Title of the project")
	genCMD.StringVar(&config.Author, "author", "", "Author of the project")
	genCMD.StringVar(&config.DbType, "db", "sqlite", "Database type")
	genCMD.StringVar(&config.Arch, "arch", "clean", "Architecture type")

	flag.Parse()

	genCMD.Parse(os.Args[2:])

	if config.Title == "" {
		config.Title = c.Prompt("Title")
	}

	if config.Author == "" {
		config.Author = c.Prompt("Author")
	}

	if config.DbType == "" {
		config.DbType = c.PromptOptions("Database type", []string{"sqlite", "mysql", "postgres"})
	}

	if config.Arch == "" {
		config.Arch = c.PromptOptions("Architecture", []string{"clean", "onion"})
	}

	// // Print project details
	c.Display("Project generated successfully!")
	c.Display("Title:", config.Title)
	c.Display("Author:", config.Author)
	c.Display("Database Type:", config.DbType)
	c.Display("Architecture:", config.Arch)
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

func (c *CLI) Display(a ...interface{}) {
	fmt.Println(a...)
}

func (c *CLI) HandleArgs() {
	switch os.Args[1] {
	case "gen":
		c.Start()
	case "--help":
		c.project.PrintHelp()
	default:
		c.InvalidArgs()
	}
}

func (c *CLI) InvalidArgs() {
	if len(os.Args) < 2 {
		c.Display("Usage: garch-go [command] [arguments]")
		return
	}

	c.Display("Unknown command:", os.Args[1])
}
