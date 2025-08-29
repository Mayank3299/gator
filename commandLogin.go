package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return errors.New("please provide the arguments for the command")
	}

	userName := cmd.Args[0]
	err := s.configFile.SetUser(userName)
	if err != nil {
		return fmt.Errorf("unable to set current_user in config: %v", err)
	}

	fmt.Printf("Username has been set")
	return nil
}
