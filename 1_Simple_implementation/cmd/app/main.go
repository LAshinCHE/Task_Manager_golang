package main

import (
	"log"
	"os"

	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/cli"
	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/service"
	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/storage"
)

func main() {
	argsWithoutProgs := os.Args[1:]

	storageJson := storage.NewStorage("storage")

	taskManagerService := service.NewService(storageJson)

	executor := cli.NewCLI(cli.Deps{Service: taskManagerService})

	err := executor.HandleCommand(argsWithoutProgs)
	if err != nil {
		log.Fatal(err)
	}
}
