package service

import (
	"flag"
	"os"

	"github.com/Rhaqim/garch-go/config"
	"github.com/Rhaqim/garch-go/internal/adapter/cli"
	"github.com/Rhaqim/garch-go/internal/app/domain"
	"github.com/Rhaqim/garch-go/internal/app/usecase"
	"github.com/Rhaqim/garch-go/internal/utils"
)

// ProjectService provides methods for project management
type ProjectService struct {
	cli cli.CLIInterface
}

// NewProjectService creates a new instance of ProjectService
func NewProjectService(cli cli.CLIInterface) usecase.ProjectUseCase {
	return &ProjectService{
		cli: cli,
	}
}

// GenerateProject generates a new project with the given configuration
func (s *ProjectService) GenerateProject(config *domain.ProjectConfig) error {
	var err error

	genCMD := flag.NewFlagSet("gen", flag.ExitOnError)

	genCMD.StringVar(&config.Title, "title", "", "Title of the project")
	genCMD.StringVar(&config.Author, "author", "", "Author of the project")
	genCMD.StringVar(&config.DbType, "db", "", "Database type")
	genCMD.StringVar(&config.Arch, "arch", "", "Architecture type")

	flag.Parse()

	genCMD.Parse(os.Args[2:])

	s.HandleArgs(config)

	// Print project details
	s.cli.Display("Project generated successfully!")
	s.cli.Display("Title:", config.Title)
	s.cli.Display("Author:", config.Author)
	s.cli.Display("Database Type:", config.DbType)
	s.cli.Display("Architecture:", config.Arch)

	return err
}

// PrintHelp implements ProjectServiceInterface.
func (s *ProjectService) Help() {
	s.cli.Display("Usage: garch-go [command] [options]")
}

// Usage implements ProjectServiceInterface.
func (s *ProjectService) Usage() {
	s.cli.Display("Usage: garch-go [command] [options]")
}

func (s *ProjectService) HandleArgs(projectConfig *domain.ProjectConfig) {
	// Select the type of the project
	projectTypes := utils.GetFields(domain.Deps)

	if projectConfig.Type == "" {
		projectConfig.Type = s.cli.PromptOptions("Type of the project", projectTypes)
	}

	// Select the architecture type
	if projectConfig.Arch == "" {
		projectConfig.Arch = s.cli.PromptOptions("Architecture", config.ArchTypes)
	}

	if projectConfig.Title == "" {
		projectConfig.Title = s.cli.Prompt("Title")
	}

	if projectConfig.Author == "" {
		projectConfig.Author = s.cli.Prompt("Author")
	}

	if projectConfig.DbType == "" {
		projectConfig.DbType = s.cli.PromptOptions("Database type", []string{"sqlite", "mysql", "postgres"})
	}
}
