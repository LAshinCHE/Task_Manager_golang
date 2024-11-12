package main

import (
	"github.com/LAshinCHE/Task_Manager_golang/internal/cli"
	"github.com/LAshinCHE/Task_Manager_golang/internal/service"
	"github.com/LAshinCHE/Task_Manager_golang/internal/storage"
)

func main() {
	JsonStorage := storage.NewStorage()

	service := service.NewService(JsonStorage)

	comandHendler := cli.NewCLI(service)
}
