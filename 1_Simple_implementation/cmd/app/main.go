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

	st := storage.NewStorage("storage")

	serv := service.NewService(st)

	err := cli.HandleCommand(argsWithoutProgs, serv)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
}
