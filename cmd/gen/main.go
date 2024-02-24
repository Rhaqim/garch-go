package main

import (
	"github.com/Rhaqim/garch-go/internal/app/core"
)

func main() {

	// Initialize application
	core := core.NewCore()

	// Run the application
	core.Run()

}
