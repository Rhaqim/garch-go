package core

import (
	"github.com/Rhaqim/garch-go/config"
	"github.com/Rhaqim/garch-go/internal/app/domain"
	arch "github.com/Rhaqim/garch-go/internal/infrastructure/architectures"
	"github.com/Rhaqim/garch-go/internal/utils"
)

type Core struct {
	Project     *domain.ProjectConfig
	Folders     []domain.FolderStructure
	Files       []domain.FileStructure
	Dependecies []string
	ProjectPath string
}

func NewCore(project *domain.ProjectConfig) CoreInterface {

	// Get the architecture layout
	architecture := arch.ArchitetureMap[config.ArchitecureType(utils.NormalizeString(project.Arch))]

	return &Core{
		Project:     project,
		Folders:     architecture.Folders,
		Files:       architecture.Files,
		Dependecies: domain.DepsStringMap[project.Type],
		ProjectPath: project.Output,
	}
}

func (c *Core) Generate() {
	// Check if the project path is set
	if c.ProjectPath != "" {
		path := utils.OutputPathHandler(c.ProjectPath)

		CreateFolder(path)
		ChangeDirectory(path)

		c._generateFolderStructure()

		return
	}
	c._generateFolderStructure()

}

func (c *Core) _generateFolderStructure() {

	CreateFolder(c.Project.Title)
	ChangeDirectory(c.Project.Title)

	// git init
	// RunGitInit()

	// go mod init
	RunGoInit(c.Project.Author, c.Project.Title)

	// Go get dependencies
	if len(c.Dependecies) > 0 {
		RunGoGet(c.Dependecies...)
	}

	// Generate the folder structure
	GenerateRecursive(c.Folders...)

	// Generate root files
	if len(c.Files) > 0 {
		for _, file := range c.Files {
			CreateFile(file.FileName, file.FileContent)
		}
	}
}
