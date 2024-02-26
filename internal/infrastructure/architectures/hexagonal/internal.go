package hexagonal

import "github.com/Rhaqim/garch-go/internal/app/domain"

// Folders: adapters, core
const ADAPTERS_FOLDER = "adapters"
const CORE_FOLDER = "core"
const CONFIG_FOLDER = "config"

// Adapters Folder
// Folders: cache, handler, repository, token
const CACHE_FOLDER = "cache"
const HANDLER_FOLDER = "handler"
const REPOSITORY_FOLDER = "repository"
const TOKEN_FOLDER = "token"

// Cache Files
const CACHE_FILE = "cache.go"

// Handler Files
const HANDLER_FILE = "handler.go"

// Repository Files
const REPOSITORY_FILE = "repository.go"

// Token Files
const TOKEN_FILE = "token.go"

// Core Folder
// Folders: domain, port, service, util
const DOMAIN_FOLDER = "domain"
const PORT_FOLDER = "port"
const SERVICE_FOLDER = "service"
const UTIL_FOLDER = "util"

// Domain Files
const ENTITY_FILE = "entity.go"
const VALUE_OBJECT_FILE = "value_object.go"

// Port Files
const INPUT_PORT_FILE = "input_port.go"
const OUTPUT_PORT_FILE = "output_port.go"

// Service Files
const SERVICE_FILE = "service.go"

// Util Files
const ERROR_FILE = "error.go"
const LOGGER_FILE = "logger.go"

// Config Files
const CONFIG_FILE = "config.go"

var HexInternalFolders []domain.FolderStructure = []domain.FolderStructure{
	{
		FolderTitle: ADAPTERS_FOLDER,
		SubFolders: []domain.FolderStructure{
			{
				FolderTitle: CACHE_FOLDER,
				Files: []domain.FileStructure{
					{
						FileName:    CACHE_FILE,
						FileContent: `package cache`,
					},
				},
			},
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
				FolderTitle: REPOSITORY_FOLDER,
				Files: []domain.FileStructure{
					{
						FileName:    REPOSITORY_FILE,
						FileContent: `package repository`,
					},
				},
			},
			{
				FolderTitle: TOKEN_FOLDER,
				Files: []domain.FileStructure{
					{
						FileName:    TOKEN_FILE,
						FileContent: `package token`,
					},
				},
			},
		},
	},
	{
		FolderTitle: CORE_FOLDER,
		SubFolders: []domain.FolderStructure{
			{
				FolderTitle: DOMAIN_FOLDER,
				Files: []domain.FileStructure{
					{
						FileName:    ENTITY_FILE,
						FileContent: `package domain`,
					},
					{
						FileName:    VALUE_OBJECT_FILE,
						FileContent: `package domain`,
					},
				},
			},
			{
				FolderTitle: PORT_FOLDER,
				Files: []domain.FileStructure{
					{
						FileName:    INPUT_PORT_FILE,
						FileContent: `package port`,
					},
					{
						FileName:    OUTPUT_PORT_FILE,
						FileContent: `package port`,
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
			{
				FolderTitle: UTIL_FOLDER,
				Files: []domain.FileStructure{
					{
						FileName:    ERROR_FILE,
						FileContent: `package util`,
					},
					{
						FileName:    LOGGER_FILE,
						FileContent: `package util`,
					},
				},
			},
		},
	},
	{
		FolderTitle: CONFIG_FOLDER,
		Files: []domain.FileStructure{
			{
				FileName:    CONFIG_FILE,
				FileContent: `package config`,
			},
		},
	},
}
