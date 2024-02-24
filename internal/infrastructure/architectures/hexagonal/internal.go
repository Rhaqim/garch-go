package architecture

import "github.com/Rhaqim/garch-go/internal/app/core"

// Folders: adapters, core
const ADAPTERS_FOLDER = "adapters"
const CORE_FOLDER = "core"

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

var HexInternalFolders []core.FolderStructure = []core.FolderStructure{
	{
		FolderTitle: ADAPTERS_FOLDER,
		Files:       []core.FileStructure{},
		SubFolders: []core.FolderStructure{
			{
				FolderTitle: CACHE_FOLDER,
				Files: []core.FileStructure{
					{
						FileName:    CACHE_FILE,
						FileContent: `package cache`,
					},
				},
				SubFolders: []core.FolderStructure{},
			},
			{
				FolderTitle: HANDLER_FOLDER,
				Files: []core.FileStructure{
					{
						FileName:    HANDLER_FILE,
						FileContent: `package handler`,
					},
				},
				SubFolders: []core.FolderStructure{},
			},
			{
				FolderTitle: REPOSITORY_FOLDER,
				Files: []core.FileStructure{
					{
						FileName:    REPOSITORY_FILE,
						FileContent: `package repository`,
					},
				},
				SubFolders: []core.FolderStructure{},
			},
			{
				FolderTitle: TOKEN_FOLDER,
				Files: []core.FileStructure{
					{
						FileName:    TOKEN_FILE,
						FileContent: `package token`,
					},
				},
				SubFolders: []core.FolderStructure{},
			},
		},
	},
	{
		FolderTitle: CORE_FOLDER,
		Files:       []core.FileStructure{},
		SubFolders: []core.FolderStructure{
			{
				FolderTitle: DOMAIN_FOLDER,
				Files: []core.FileStructure{
					{
						FileName:    ENTITY_FILE,
						FileContent: `package domain`,
					},
					{
						FileName:    VALUE_OBJECT_FILE,
						FileContent: `package domain`,
					},
				},
				SubFolders: []core.FolderStructure{},
			},
			{
				FolderTitle: PORT_FOLDER,
				Files: []core.FileStructure{
					{
						FileName:    INPUT_PORT_FILE,
						FileContent: `package port`,
					},
					{
						FileName:    OUTPUT_PORT_FILE,
						FileContent: `package port`,
					},
				},
				SubFolders: []core.FolderStructure{},
			},
			{
				FolderTitle: SERVICE_FOLDER,
				Files: []core.FileStructure{
					{
						FileName:    SERVICE_FILE,
						FileContent: `package service`,
					},
				},
				SubFolders: []core.FolderStructure{},
			},
			{
				FolderTitle: UTIL_FOLDER,
				Files: []core.FileStructure{
					{
						FileName:    ERROR_FILE,
						FileContent: `package util`,
					},
					{
						FileName:    LOGGER_FILE,
						FileContent: `package util`,
					},
				},
				SubFolders: []core.FolderStructure{},
			},
		},
	},
}
