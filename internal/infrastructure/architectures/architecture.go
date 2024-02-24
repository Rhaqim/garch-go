package architectures

import (
	"github.com/Rhaqim/garch-go/config"
	"github.com/Rhaqim/garch-go/internal/app/domain"
	"github.com/Rhaqim/garch-go/internal/infrastructure/architectures/hexagonal"
)

var ArchitetureMap = map[config.ArchitecureType]domain.ArchitectureLayout{
	config.Clean: {},
	config.Onion: {},
	config.Hexagonal: {
		Files:   hexagonal.HexRootFiles,
		Folders: hexagonal.HexRootFolders,
	},
}
