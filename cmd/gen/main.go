package main

import (
	"flag"
	"os"

	"github.com/Rhaqim/garch-go/internal/adapter/cli"
	"github.com/Rhaqim/garch-go/internal/app"
)

func main() {
	// Parse command-line flags
	var title, author, dbType, arch string

	genCMD := flag.NewFlagSet("gen", flag.ExitOnError)

	genCMD.StringVar(&title, "title", "", "Title of the project")
	genCMD.StringVar(&author, "author", "", "Author of the project")
	genCMD.StringVar(&dbType, "db", "sqlite", "Database type")
	genCMD.StringVar(&arch, "arch", "clean", "Architecture type")

	flag.Parse()

	// Initialize CLI adapter
	cli := cli.NewCLI()

	if len(os.Args) < 2 {
		cli.Display("Usage: garch-go [command] [arguments]")
		return
	}

	// Initialize application service
	projectService := app.NewProjectService()

	switch os.Args[1] {
	case "gen":
		genCMD.Parse(os.Args[2:])
	case "--help":
		cli.Display("Usage: garch-go [command] [arguments]")
	default:
		cli.Display("Unknown command:", os.Args[1])
		return
	}

	// genCMD.Parse(flag.Args())

	// if not arguments not provided, provide prompt
	if title == "" {
		title = cli.Prompt("Title")
	}

	// Generate project
	project, err := projectService.GenerateProject(title, author, dbType, arch)
	if err != nil {
		cli.Display("Error generating project:", err)
		return
	}

	// Print project details
	cli.Display("Project generated successfully!")
	cli.Display("Title:", project.Title)
	cli.Display("Author:", project.Author)
	cli.Display("Database Type:", project.DbType)
	cli.Display("Architecture:", project.Arch)
}
