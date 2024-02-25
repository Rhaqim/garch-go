package domain

// ProjectConfig represents the configuration for a project
type ProjectConfig struct {
	Type   string `short:"t" long:"type" description:"Type of the project"`
	Arch   string `short:"a" long:"architecture" description:"Architecture type"`
	Title  string `short:"n" long:"name" description:"Title of the project"`
	Author string `short:"u" long:"user" description:"Author of the project"`
	DbType string `short:"d" long:"db" description:"Database type"`
	// Add more fields for other configurations
}

type ProjectDependencies struct {
	API     []string
	CLI     []string
	Lib     []string
	Service []string
	Other   []string
}

type FolderStructure struct {
	FolderTitle string
	SubFolders  []FolderStructure
	Files       []FileStructure
}

type FileStructure struct {
	FileName    string
	FileContent string
}

type ArchitectureLayout struct {
	Files   []FileStructure
	Folders []FolderStructure
}
