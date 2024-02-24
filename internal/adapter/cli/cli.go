package cli

import (
	"fmt"
)

// CLI represents the CLI adapter
type CLI struct{}

// NewCLI creates a new instance of CLI
func NewCLI() CLIInterface {
	return &CLI{}
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
