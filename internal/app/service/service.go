package service

import (
	"flag"
	"os"
	"reflect"

	"github.com/Rhaqim/garch-go/config"
	"github.com/Rhaqim/garch-go/internal/adapter/cli"
	"github.com/Rhaqim/garch-go/internal/app/core"
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

	genCMD := Parser(config)

	genCMD.Parse(os.Args[2:])

	s.HandleArgs(config)

	new_core := core.NewCore(config)

	// Generate the project
	new_core.Generate()

	// Print project details
	s.cli.Display("Project generated successfully!")
	s.cli.Display("Project Type:", config.Type)
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
	defaultUsername := utils.GetGitUsername()
	defaultTitle := "MyProject"
	defaultType := "Other"
	defaultArch := "Hexagonal"
	defaultDB := "sqlite"

	// Select the type of the project
	projectTypes := utils.GetFields(domain.Deps)

	if projectConfig.Type == "" {
		projectConfig.Type = s.cli.PromptOptions("What type of project are you building? default: "+defaultType, projectTypes)
	}

	// Select the architecture type
	if projectConfig.Arch == "" {
		projectConfig.Arch = s.cli.PromptOptions("Select the architecture type default: "+defaultArch, config.ArchTypes)
	}

	if projectConfig.Title == "" {
		projectConfig.Title = s.cli.Prompt("Project title default: "+defaultTitle, defaultTitle)
	}

	if projectConfig.Author == "" {
		projectConfig.Author = s.cli.Prompt("Author default: "+defaultUsername, defaultUsername)
	}

	if projectConfig.DbType == "" {
		ok := s.cli.Bool("Do you want to use a database?")
		if !ok {
			return
		}
		projectConfig.DbType = s.cli.PromptOptions("Database type default: "+defaultDB, []string{"sqlite", "mysql", "postgres", "mssql"})
	}
}

// Parser parses the project configuration
// and returns a flagset
// It works by using reflection to get the fields of the struct
// and then adding the fields to the flagset
func Parser(config *domain.ProjectConfig) *flag.FlagSet {
	genCMD := flag.NewFlagSet("gen", flag.ExitOnError)

	t := reflect.TypeOf(config)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// get the tags of the field
		tag := field.Tag
		short := tag.Get("short")
		long := tag.Get("long")
		description := tag.Get("description")

		genCMD.StringVar(GetField(config, field.Name), short, "", description)
		genCMD.StringVar(GetField(config, field.Name), long, "", description)

	}

	return genCMD
}

// GetField returns the field of a struct
// It works by using reflection to get the field of the struct
// and then returning the field
func GetField(config *domain.ProjectConfig, fieldName string) *string {
	v := reflect.ValueOf(config)
	f := reflect.Indirect(v).FieldByName(fieldName)
	return f.Addr().Interface().(*string)
}
