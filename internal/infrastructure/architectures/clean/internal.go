package clean

import (
	"github.com/Rhaqim/garch-go/internal/app/domain"
	repo "github.com/Rhaqim/garch-go/internal/infrastructure/repository"
)

// Folders: app, delivery, infrastructure
const APP_FOLDER = "app"
const DELIVERY_FOLDER = "delivery"
const INFRASTRUCTURE_FOLDER = "infrastructure"

// Delivery Folder
// Folders: handler, presenter
const HANDLER_FOLDER = "handler"
const PRESENTER_FOLDER = "presenter"

// Handler Files
const HANDLER_FILE = "handler.go"

// Presenter Files
const PRESENTER_FILE = "presenter.go"

var CleanInternalFolders []domain.FolderStructure = []domain.FolderStructure{
	{
		FolderTitle: APP_FOLDER,
		SubFolders:  repo.AppFolders,
	},
	{
		FolderTitle: DELIVERY_FOLDER,
		SubFolders: []domain.FolderStructure{
			{
				FolderTitle: HANDLER_FOLDER,
				Files: []domain.FileStructure{
					{
						FileName:    HANDLER_FILE,
						FileContent: `package handler`,
					},
				},
			},
			{
				FolderTitle: PRESENTER_FOLDER,
				Files: []domain.FileStructure{
					{
						FileName:    PRESENTER_FILE,
						FileContent: `package presenter`,
					},
				},
			},
		},
	},
	{
		FolderTitle: INFRASTRUCTURE_FOLDER,
		SubFolders:  repo.InfrastructureFolders,
	},
}
