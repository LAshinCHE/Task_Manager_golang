package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/LAshinCHE/Task_Manager_golang/internal/cli"
	"github.com/LAshinCHE/Task_Manager_golang/internal/service"
	"github.com/LAshinCHE/Task_Manager_golang/internal/storage"
)

func main() {
	JsonStorage := storage.NewStorage()

	taskManagerService := service.NewService(JsonStorage)

	comandHendler := cli.NewCLI(taskManagerService)

	scan := bufio.NewScanner(os.Stdin)
	fmt.Println("___Welcome to the Personal Task Manager___\nA program for recording tasks.\nTo read the instructions, enter /help")
	for {
		fmt.Println("> ")

		scan.Scan()
		command := strings.TrimSpace(scan.Text())

		if command == "quit" {
			fmt.Println("The program considered the quit command. The program has completed its execution")
			return
		}

		args := strings.Fields(command)

		if err := comandHendler.Execute(args); err != nil {
			fmt.Println("Error while running: ", err)
			os.Exit(1)
		}
	}
}
