package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/app/comands"
)

func main() {
	argsWithoutProgs := os.Args[1:]

	c := comands.NewCommander(argsWithoutProgs)

	if len(argsWithoutProgs) == 0 {
		c.Help()
		log.Fatal("The program is running without arguments, please enter the arguments from the list")
	}
	fmt.Print("Hello golang")
}
