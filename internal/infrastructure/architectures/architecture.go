package architectures

import (
	"github.com/Rhaqim/garch-go/config"
	"github.com/Rhaqim/garch-go/internal/app/domain"
	"github.com/Rhaqim/garch-go/internal/infrastructure/architectures/clean"
	"github.com/Rhaqim/garch-go/internal/infrastructure/architectures/hexagonal"
	"github.com/Rhaqim/garch-go/internal/infrastructure/architectures/onion"
)

var ArchitetureMap = map[config.ArchitecureType]domain.ArchitectureLayout{
	config.Clean: {
		Files:   clean.CleanRootFiles,
		Folders: clean.CleanRootFolders,
	},
	config.Onion: {
		Files:   onion.OnionRootFiles,
		Folders: onion.OnionRootFolders,
	},
	config.Hexagonal: {
		Files:   hexagonal.HexRootFiles,
		Folders: hexagonal.HexRootFolders,
	},
	config.DDD: {},
}
