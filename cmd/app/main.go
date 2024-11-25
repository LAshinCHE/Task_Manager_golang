package main

import (
	"context"
	"fmt"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/LAshinCHE/Task_Manager_golang/internal/cli"
	"github.com/LAshinCHE/Task_Manager_golang/internal/service"
	"github.com/LAshinCHE/Task_Manager_golang/internal/storage"
)

const (
	fileName = "task_storage.json"
)

func main() {
	var (
		ctx                   = context.Background()
		ctxWithCancel, cancel = signal.NotifyContext(ctx, syscall.SIGINT)
		wg                    sync.WaitGroup
		commandChan           = make(chan []string)
		errorChan             = make(chan error)
	)

	defer cancel()

	JsonStorage := storage.NewStorage(fileName)
	taskManagerService := service.NewService(JsonStorage)
	commandHandler := cli.NewCLI(taskManagerService)

	// Запуск горутины для обработки ввода команд
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(commandChan)
		for {
			select {
			case <-ctxWithCancel.Done():
				fmt.Println("Graceful shutdown")
				return
			default:
				fmt.Print("Enter command: ")
				var input string
				// Считывание строки без блокировки сигнала
				_, err := fmt.Scanln(&input)
				if err != nil {
					if strings.Contains(err.Error(), "unexpected newline") {
						continue
					}
					fmt.Println("Error reading input:", err)
					cancel()
					return
				}

				command := strings.TrimSpace(input)
				if command == "quit" || len(command) == 0 {
					cancel()
					fmt.Println("The program considered the quit command. The program has completed its execution")
					return
				}
				args := strings.Fields(command)
				commandChan <- args
			}
		}
	}()

	// Горутинa для выполнения команд
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(errorChan)
		for {
			select {
			case <-ctxWithCancel.Done():
				fmt.Println("Command processing goroutine ended")
				return
			case command := <-commandChan:
				errorChan <- commandHandler.Execute(command)
			}
		}
	}()

	// Горутинa для обработки ошибок
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctxWithCancel.Done():
				fmt.Println("Error handling goroutine ended")
				return
			case err := <-errorChan:
				if err != nil {
					fmt.Println(err)
					cancel()
				}
			}
		}
	}()

	// Ожидание завершения всех процессов
	wg.Wait()
}
