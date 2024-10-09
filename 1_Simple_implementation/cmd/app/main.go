package main

import (
	"log"
	"os"

	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/service"
	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/storage"
)

func main() {
	argsWithoutProgs := os.Args[1:]

	filename := "Storage_1"

	st := storage.NewStorage(filename)
	s := service.NewService(st)

	if len(argsWithoutProgs) == 0 {
		log.Fatal("The program is running without arguments, please enter the arguments from the list")
	}

}
