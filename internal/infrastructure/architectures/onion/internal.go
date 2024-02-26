package onion

import (
	"github.com/Rhaqim/garch-go/internal/app/domain"
	repo "github.com/Rhaqim/garch-go/internal/infrastructure/repository"
)

// Folders: app, interface, infrastructure
const APP_FOLDER = "app"
const INTERFACE_FOLDER = "interface"
const INFRASTRUCTURE_FOLDER = "infrastructure"

// Interface Folder
// Folders: handler, presenter
const DELIVERY_FOLDER = "delivery"
const PERSISTENCE_FOLDER = "persistance"

// Handler Files
const DELIVERY_FILE = "delivery.go"

var OnionInternalFolders []domain.FolderStructure = []domain.FolderStructure{
	{
		FolderTitle: APP_FOLDER,
		SubFolders:  repo.AppFolders,
	},
	{
		FolderTitle: INTERFACE_FOLDER,
		SubFolders:  repo.InterfaceFolders,
	},
	{
		FolderTitle: INFRASTRUCTURE_FOLDER,
		SubFolders:  repo.InfrastructureFolders,
	},
}
