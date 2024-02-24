package core

import (
	"os"

	"github.com/Rhaqim/garch-go/internal/adapter/cli"
	"github.com/Rhaqim/garch-go/internal/app/domain"
	"github.com/Rhaqim/garch-go/internal/app/usecase"
)

type Core struct {
	project usecase.ProjectUseCase
}

func NewCore(project usecase.ProjectUseCase) *Core {
	return &Core{
		project: project,
	}
}

func (c *Core) Run() {
	// Initialize CLI
	var cli cli.CLIInterface = cli.NewCLI()

	// Parse the command line arguments
	config := domain.ProjectConfig{}

	cli.Display("Welcome to Garch! \n")

	if len(os.Args) < 2 {
		cli.InvalidArgs()
		return
	}

	switch os.Args[1] {
	case "gen":
		cli.Start(&config)
	case "--help":
		cli.Usage()
	default:
		cli.InvalidArgs()
	}

	// Print project details
	cli.Display("Project generated successfully!")
	cli.Display("Title:", config.Title)
	cli.Display("Author:", config.Author)
	cli.Display("Database Type:", config.DbType)
	cli.Display("Architecture:", config.Arch)

}
