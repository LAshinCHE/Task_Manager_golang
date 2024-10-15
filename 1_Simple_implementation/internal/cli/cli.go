package cli

import (
	"fmt"

	"github.com/LAshinCHE/Task_Manager_golang/1_Simple_implementation/internal/service"
)

func HandleCommand(commandWithArg []string, s *service.Service) error {
	var command string
	if len(commandWithArg) == 0 {
		return fmt.Errorf("Enter -help to check list command u can write to this program")
	} else {
		command = commandWithArg[0]
	}
	switch command {
	case "-help":
		s.Help()
	default:
		fmt.Println("Chose your commands from list")
		s.Help()
	}

	return nil
}
