package main

import (
	"github.com/Rhaqim/garch-go/internal/app/core"
	"github.com/Rhaqim/garch-go/internal/app/service"
)

func main() {

	// Initialize project service
	project := service.NewProjectService()

	// Initialize application
	core := core.NewCore(project)

	// Run the application
	core.Run()

}
