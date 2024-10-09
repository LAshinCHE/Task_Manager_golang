package main

import (
	"log"
	"os"

	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/commander"
	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/service"
	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/storage"
)

func main() {
	argsWithoutProgs := os.Args[1:]

	filename := "Storage_1"

	st := storage.NewStorage(filename)
	serv := service.NewService(st)

	err := commander.HandleCommand(argsWithoutProgs, serv)
	if err != nil {
		log.Fatal(err)
	}
}
