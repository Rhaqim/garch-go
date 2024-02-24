package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/Rhaqim/garch-go/config"
	"github.com/Rhaqim/garch-go/internal/app/domain"
	"github.com/Rhaqim/garch-go/internal/utils"
)

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

	c.HandleArgs(config)

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

func (c *CLI) HandleArgs(projectConfig *domain.ProjectConfig) {
	// Select the type of the project
	projectTypes := utils.GetFields(domain.Deps)

	if projectConfig.Type == "" {
		projectConfig.Type = c.PromptOptions("Type of the project", projectTypes)
	}

	// Select the architecture type
	if projectConfig.Arch == "" {
		projectConfig.Arch = c.PromptOptions("Architecture", config.ArchTypes)
	}

	if projectConfig.Title == "" {
		projectConfig.Title = c.Prompt("Title")
	}

	if projectConfig.Author == "" {
		projectConfig.Author = c.Prompt("Author")
	}

	if projectConfig.DbType == "" {
		projectConfig.DbType = c.PromptOptions("Database type", []string{"sqlite", "mysql", "postgres"})
	}
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
