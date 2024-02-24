package core

type FolderStructure struct {
	FolderTitle string
	SubFolders  []FolderStructure
	Files       []FileStructure
}

type FileStructure struct {
	FileName    string
	FileContent string
}

type CoreInterface interface {
	Generate()
}
