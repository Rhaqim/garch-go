package core

import (
	"github.com/Rhaqim/garch-go/internal/app/domain"
	hex "github.com/Rhaqim/garch-go/internal/infrastructure/architectures/hexagonal"
)

type Core struct {
	Project *domain.ProjectConfig
	Folders []domain.FolderStructure
	Files   []domain.FileStructure
}

func NewCore(project *domain.ProjectConfig) CoreInterface {
	folders := hex.HexRootFolders
	files := hex.HexRootFiles

	return &Core{
		Project: project,
		Folders: folders,
		Files:   files,
	}
}

func (c *Core) Generate() {
	CreateFolder(c.Project.Title)
	ChangeDirectory(c.Project.Title)

	// git init
	// RunGitInit()

	// go mod init
	RunGoInit(c.Project.Author, c.Project.Title)

	// Generate the folder structure
	GenerateRecursive(c.Folders...)

	// Generate root files
	if len(c.Files) > 0 {
		for _, file := range c.Files {
			CreateFile(file.FileName, file.FileContent)
		}
	}
}
