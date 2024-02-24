package core

import "github.com/Rhaqim/garch-go/internal/app/domain"

type Core struct {
	Project *domain.ProjectConfig
	Folders []FolderStructure
	Files   []FileStructure
}

func NewCore(project *domain.ProjectConfig) CoreInterface {
	return &Core{
		Project: project,
	}
}

func (c *Core) Generate() {
	CreateFolder(c.Project.Title)
	ChangeDirectory(c.Project.Title)

	// git init
	RunGitInit()

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
