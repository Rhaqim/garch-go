package ddd

import (
	"github.com/Rhaqim/garch-go/internal/app/domain"
	repo "github.com/Rhaqim/garch-go/internal/infrastructure/repository"
)

const DOMAIN_FOLDER = "domain"
const APPLICATION_FOLDER = "application"
const INTERFACE_FOLDER = "interface"
const INFRASTRUCTURE_FOLDER = "infrastructure"

// Domain folders: entity, repository, service
const ENTITY_FOLDER = "entity"
const REPOSITORY_FOLDER = "repository"
const SERVICE_FOLDER = "service"

// Entity Files
const ENTITY_FILE = "entity.go"

// Repository Files
const REPOSITORY_FILE = "repository.go"

// Service Files
const SERVICE_FILE = "service.go"

// Application folders: service, dto
const SERVICE_FOLDER_APP = "service"
const DTO_FOLDER = "dto"

// Service Files
const SERVICE_FILE_APP = "service.go"

// Dto Files
const DTO_FILE = "dto.go"

// Internal folder for ddd architecture
var DDDInternalFolders []domain.FolderStructure = []domain.FolderStructure{
	{
		FolderTitle: DOMAIN_FOLDER,
		SubFolders: []domain.FolderStructure{
			{
				FolderTitle: ENTITY_FOLDER,
				Files: []domain.FileStructure{
					{
						FileName:    ENTITY_FILE,
						FileContent: `package entity`,
					},
				},
			},
			{
				FolderTitle: REPOSITORY_FOLDER,
				Files: []domain.FileStructure{
					{
						FileName:    REPOSITORY_FILE,
						FileContent: `package repository`,
					},
				},
			},
			{
				FolderTitle: SERVICE_FOLDER,
				Files: []domain.FileStructure{
					{
						FileName:    SERVICE_FILE,
						FileContent: `package service`,
					},
				},
			},
		},
	},
	{
		FolderTitle: APPLICATION_FOLDER,
		SubFolders: []domain.FolderStructure{
			{
				FolderTitle: SERVICE_FOLDER_APP,
				Files: []domain.FileStructure{
					{
						FileName:    SERVICE_FILE_APP,
						FileContent: `package service`,
					},
				},
			},
			{
				FolderTitle: DTO_FOLDER,
				Files: []domain.FileStructure{
					{
						FileName:    DTO_FILE,
						FileContent: `package dto`,
					},
				},
			},
		},
	},
	{
		FolderTitle: INFRASTRUCTURE_FOLDER,
		SubFolders:  repo.InfrastructureFolders,
	},
	{
		FolderTitle: INTERFACE_FOLDER,
		SubFolders:  repo.InterfaceFolders,
	},
}
