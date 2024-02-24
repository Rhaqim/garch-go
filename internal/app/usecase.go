package app

// ProjectUseCase defines use cases for project management
type ProjectUseCase interface {
	GenerateProject(title, author, dbType, arch string) (*ProjectConfig, error)
	PrintHelp()
}
