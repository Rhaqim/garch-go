package usecase

import "github.com/Rhaqim/garch-go/internal/app/domain"

// ProjectUseCase defines use cases for project management
type ProjectUseCase interface {
	// GenerateProject generates a new project
	// with the given title, author, dbType, and arch
	GenerateProject(config *domain.ProjectConfig) error
	// PrintHelp prints the help message
	Help()
	// List lists the available project types
	List(item ...string)
	// HandleArgs handles the command-line arguments
	HandleArgs(config *domain.ProjectConfig)
}
