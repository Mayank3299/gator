package main

import (
	"context"
	"fmt"

	"github.com/Mayank3299/BlogAggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		currentUser, err := s.db.GetUser(context.Background(), s.configFile.CurrentUser)
		if err != nil {
			return fmt.Errorf("user not found - %v", err)
		}

		return handler(s, cmd, currentUser)
	}
}
