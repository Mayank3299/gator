package main

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return errors.New("please provide the arguments for the command")
	}

	userName := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		os.Exit(1)
		return fmt.Errorf("user not present in db -%v", err)
	}

	err = s.configFile.SetUser(userName)
	if err != nil {
		return fmt.Errorf("unable to set current_user in config: %v", err)
	}

	fmt.Println("Username has been set")
	return nil
}
