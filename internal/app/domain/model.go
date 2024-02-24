package domain

// ProjectConfig represents the configuration for a project
type ProjectConfig struct {
	Type   string `short:"t" long:"type" description:"Type of the project"`
	Arch   string `short:"a" long:"arch" description:"Architecture type"`
	Title  string `short:"t" long:"title" description:"Title of the project"`
	Author string `short:"u" long:"author" description:"Author of the project"`
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

var Deps = ProjectDependencies{
	API:     []string{"echo", "gin", "fiber", "chi"},
	CLI:     []string{"cobra", "cli"},
	Lib:     []string{"ginkgo", "gomega", "testify"},
	Service: []string{"grpc", "nats", "rabbitmq"},
	Other:   []string{},
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
