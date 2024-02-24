package core

import (
	"os"
	"os/exec"
)

func CreateFolder(title string) {
	// Create a new directory
	err := os.Mkdir(title, 0755)
	if err != nil {
		panic(err)
	}
}

func ChangeDirectory(title string) {
	// Change the working directory
	err := os.Chdir(title)
	if err != nil {
		panic(err)
	}
}

func CreateFile(title string, content string) {
	// Create a new file
	file, err := os.Create(title)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write
	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}

func GenerateRecursive(items ...FolderStructure) {
	// Generate the folder structure
	if len(items) > 0 {
		for _, item := range items {
			CreateFolder(item.FolderTitle)
			ChangeDirectory(item.FolderTitle)

			if len(item.Files) > 0 {
				for _, file := range item.Files {
					CreateFile(file.FileName, file.FileContent)
				}
			}

			GenerateRecursive(item.SubFolders...)
			ChangeDirectory("..")
		}
	}

}

func RunGitInit() {
	// Run git init in the current directory
	cmd := exec.Command("git", "init", "&&", "git", "branch", "-m", "main")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func RunGoInit(username, projectTitle string) {

	// Run go mod init in the current directory
	cmd := exec.Command("go", "mod", "init", "github.com/"+username+"/"+projectTitle)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
