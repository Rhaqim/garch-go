package main

import (
	"github.com/Rhaqim/garch-go/internal/adapter/cli"
	"github.com/Rhaqim/garch-go/internal/app/core"
	"github.com/Rhaqim/garch-go/internal/app/service"
)

func main() {

	// Initialize CLI
	cli := cli.NewCLI()

	// Initialize project service
	project := service.NewProjectService()

	// Initialize application
	core := core.NewCore(project, cli)

	// Run the application
	core.Run()

}
