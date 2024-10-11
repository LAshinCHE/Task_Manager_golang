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

	st, err := storage.NewStorage()

	if err != nil {
		log.Fatal(err)
	}

	serv := service.NewService(st)

	err = commander.HandleCommand(argsWithoutProgs, serv)
	if err != nil {
		log.Fatal(err)
	}

	err = st.CloseStorage()
	if err != nil {
		log.Fatal(err)
	}
}
