package main

import (
	"os"

	"github.com/Rhaqim/garch-go/internal/adapter/cli"
	"github.com/Rhaqim/garch-go/internal/app/domain"
	"github.com/Rhaqim/garch-go/internal/app/service"
)

func main() {

	// Initialize CLI
	var cli cli.CLIInterface = cli.NewCLI()

	// Initialize project service
	project := service.NewProjectService(cli)

	// Parse the command line arguments
	config := domain.ProjectConfig{}

	cli.Display("Welcome to Garch! \n")

	if len(os.Args) < 2 {
		project.Help()
		return
	}

	switch os.Args[1] {
	case "gen":
		project.GenerateProject(&config)
	case "--help", "-h":
		project.Help()
	case "--list", "-l":
		project.List()
	case "--list-type", "-lt":
		project.List("type")
	case "--list-arch", "-la":
		project.List("arch")
	default:
		project.Help()
	}

}
