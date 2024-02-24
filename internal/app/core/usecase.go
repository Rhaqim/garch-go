package core

type FolderStructure struct {
	FolderTitle string
	SubFolders  []FolderStructure
	Files       []FileStructure
}

type FileStructure struct {
	FileName string
	Content  string
}

type CoreInterface interface {
	Generate()
}
