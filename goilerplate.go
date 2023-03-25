package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/akamensky/argparse"
)

func main() {

	parser := argparse.NewParser("goilerplate", "Create a new go project")
	projectName := parser.String("n", "name", &argparse.Options{Required: true, Help: "Name of the project"})
	projectPath := parser.String("p", "path", &argparse.Options{Required: true, Help: "Path to the project directory"})
	overwriteIfExists := parser.Flag("o", "overwrite", &argparse.Options{Required: false, Help: "Overwrite the project if it already exists"})

	err := parser.Parse(os.Args)

	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	completePath := filepath.Join(*projectPath, *projectName)

	if !*overwriteIfExists {
		if _, err := os.Stat(completePath); !os.IsNotExist(err) {
			fmt.Println("Directory already exists. If you wish to overwrite it, use the -o flag")
			os.Exit(1)
		}
	}
	fmt.Println("Creating project:", *projectName, "at path:", *projectPath)

	mkdirCmd := exec.Command("mkdir", completePath)
	mkdirCmd.Run()
	// create the project file by copying the template
	// get expanded ~/
	home, _ := os.UserHomeDir()
	templatePath := filepath.Join(home, "bin/GoilerPlate/template.go")
	cpCmd := exec.Command("cp", templatePath, completePath+"/main.go")
	cpCmd.Run()

	// initialize the go mod
	initCmd := exec.Command("go", "mod", "init", *projectName)
	initCmd.Dir = completePath
	initCmd.Run()
}
