package architecture

import (
	"github.com/Rhaqim/garch-go/internal/app/domain"
	architecture "github.com/Rhaqim/garch-go/internal/infrastructure/architectures"
)

// Root files for hexagonal architecture
var HexRootFiles []domain.FileStructure = architecture.RootFiles

// RootFolders for hexagonal architecture
const CMD_FOLDER = "cmd"
const INTERNAL_FOLDER = "internal"
const PKG_FOLDER = "pkg"
const TESTDATA_FOLDER = "testdata"

var HexRootFolders []domain.FolderStructure = []domain.FolderStructure{
	{
		FolderTitle: CMD_FOLDER,
		Files: []domain.FileStructure{
			architecture.MainFile,
		},
	},
	{
		FolderTitle: INTERNAL_FOLDER,
		SubFolders:  HexInternalFolders,
	},
	{
		FolderTitle: PKG_FOLDER,
	},
	{
		FolderTitle: TESTDATA_FOLDER,
	},
}

// Internal files for hexagonal architecture
