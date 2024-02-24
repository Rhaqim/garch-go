package service

import (
	"github.com/Rhaqim/garch-go/internal/app/domain"
	"github.com/Rhaqim/garch-go/internal/app/usecase"
)

// ProjectService provides methods for project management
type ProjectService struct{}

// NewProjectService creates a new instance of ProjectService
func NewProjectService() usecase.ProjectUseCase {
	return &ProjectService{}
}

// GenerateProject generates a new project with the given configuration
func (s *ProjectService) GenerateProject(config *domain.ProjectConfig) error {
	panic("unimplemented")
}

// PrintHelp implements ProjectServiceInterface.
func (s *ProjectService) PrintHelp() {
	panic("unimplemented")
}
