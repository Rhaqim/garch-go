package onion

import (
	"github.com/Rhaqim/garch-go/internal/app/domain"
	repo "github.com/Rhaqim/garch-go/internal/infrastructure/repository"
)

// Root files for onion architecture
var OnionRootFiles []domain.FileStructure = repo.RootFiles

// RootFolders for onion architecture
const CMD_FOLDER = "cmd"
const INTERNAL_FOLDER = "internal"
const PKG_FOLDER = "pkg"
const CONFIG_FOLDER = "config"

var OnionRootFolders []domain.FolderStructure = []domain.FolderStructure{
	{
		FolderTitle: CMD_FOLDER,
		Files: []domain.FileStructure{
			repo.MainFile,
		},
	},
	{
		FolderTitle: INTERNAL_FOLDER,
		SubFolders:  OnionInternalFolders,
	},
	{
		FolderTitle: PKG_FOLDER,
	},
	{
		FolderTitle: CONFIG_FOLDER,
	},
}
