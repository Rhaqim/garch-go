package core

import (
	"os"

	"github.com/Rhaqim/garch-go/internal/adapter/cli"
	"github.com/Rhaqim/garch-go/internal/app/domain"
	"github.com/Rhaqim/garch-go/internal/app/usecase"
)

type CoreInterface interface {
	Run()
}

type Core struct {
	project usecase.ProjectUseCase
	cli     cli.CLIInterface
}

func NewCore(project usecase.ProjectUseCase, cli cli.CLIInterface) CoreInterface {
	return &Core{
		project: project,
		cli:     cli,
	}
}

func (c *Core) Run() {
	c.cli.Display("Welcome to Garch! \n")

	// Parse the command line arguments
	config := (domain.ProjectConfig{})

	if len(os.Args) < 2 {
		c.cli.InvalidArgs()
		return
	}

	switch os.Args[1] {
	case "gen":
		c.cli.Start(&config)
	case "--help":
		c.cli.Usage()
	default:
		c.cli.InvalidArgs()
	}

	// Print project details
	c.cli.Display("Project generated successfully!")
	c.cli.Display("Title:", config.Title)
	c.cli.Display("Author:", config.Author)
	c.cli.Display("Database Type:", config.DbType)
	c.cli.Display("Architecture:", config.Arch)

}
