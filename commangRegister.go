package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Mayank3299/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return errors.New("please provide the arguments for the command")
	}

	userName := cmd.Args[0]

	queryArgs := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	}

	user, err := s.db.CreateUser(context.Background(), queryArgs)
	if err != nil {
		os.Exit(1)
		return fmt.Errorf("user not created -%v", err)
	}

	s.configFile.SetUser(userName)
	fmt.Println("User has been created")
	printUser(user)

	return nil
}
