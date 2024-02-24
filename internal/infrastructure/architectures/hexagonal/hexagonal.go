package architecture

import (
	"github.com/Rhaqim/garch-go/internal/app/core"
	architecture "github.com/Rhaqim/garch-go/internal/infrastructure/architectures"
)

// Root files for hexagonal architecture
var HexRootFiles []core.FileStructure = architecture.RootFiles

// RootFolders for hexagonal architecture
const CMD_FOLDER = "cmd"
const INTERNAL_FOLDER = "internal"
const PKG_FOLDER = "pkg"
const TESTDATA_FOLDER = "testdata"

var HexRootFolders []core.FolderStructure = []core.FolderStructure{
	{
		FolderTitle: CMD_FOLDER,
		Files: []core.FileStructure{
			architecture.MainFile,
		},
		SubFolders: []core.FolderStructure{},
	},
	{
		FolderTitle: INTERNAL_FOLDER,
		Files:       []core.FileStructure{},
		SubFolders:  HexInternalFolders,
	},
	{
		FolderTitle: PKG_FOLDER,
		Files:       []core.FileStructure{},
		SubFolders:  []core.FolderStructure{},
	},
	{
		FolderTitle: TESTDATA_FOLDER,
		Files:       []core.FileStructure{},
		SubFolders:  []core.FolderStructure{},
	},
}

// Internal files for hexagonal architecture
