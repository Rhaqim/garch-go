package main

import (
	"github.com/Rhaqim/garch-go/internal/adapter/cli"
	"github.com/Rhaqim/garch-go/internal/app/service"
)

func main() {

	// Initialize application service
	projectService := service.NewProjectService()

	// Initialize CLI adapter
	cli := cli.NewCLI(projectService)

	cli.Start()

}
