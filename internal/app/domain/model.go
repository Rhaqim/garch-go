package domain

// ProjectConfig represents the configuration for a project
type ProjectConfig struct {
	Arch   string
	Title  string
	Author string
	DbType string
	// Add more fields for other configurations
}

// GarchCommands represents the available commands for the CLI application
type GarchCommands struct {
	Gen  GarchArgs `command:"gen" description:"Generate default Go project"`
	Help func()    `command:"help" description:"Display help message"`
}

// GarchArgs represents the arguments for the 'gen' command
type GarchArgs struct {
	ProjectConfig
}
