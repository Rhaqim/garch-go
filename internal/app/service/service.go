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
	cmd *flag.FlagSet
}

// NewProjectService creates a new instance of ProjectService
func NewProjectService(cli cli.CLIInterface) usecase.ProjectUseCase {
	genCMD := flag.NewFlagSet("gen", flag.ExitOnError)

	return &ProjectService{
		cli: cli,
		cmd: genCMD,
	}
}

// GenerateProject generates a new project with the given configuration
func (s *ProjectService) GenerateProject(config *domain.ProjectConfig) error {
	var err error

	// Handle the arguments
	s.HandleArgs(config)

	// Parse the project configuration
	s.Parser(config)

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
	s.cli.Display("Commands:")
	s.cli.Display("  gen [options]  Generate a new project")
	s.cli.Display("  help           Display help")
	s.Parser(&domain.ProjectConfig{})
	s.cli.Display("Options:")
	s.cmd.PrintDefaults()
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
func (s *ProjectService) Parser(config *domain.ProjectConfig) {

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

		s.cmd.StringVar(GetField(config, field.Name), short, "", description)
		s.cmd.StringVar(GetField(config, field.Name), long, "", description)

	}

	s.cmd.Parse(os.Args[2:])
}

// GetField returns the field of a struct
// It works by using reflection to get the field of the struct
// and then returning the field
func GetField(config *domain.ProjectConfig, fieldName string) *string {
	v := reflect.ValueOf(config)
	f := reflect.Indirect(v).FieldByName(fieldName)
	return f.Addr().Interface().(*string)
}
