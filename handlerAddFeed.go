package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Mayank3299/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, currentUser database.User) error {
	if len(cmd.Args) < 2 {
		return errors.New("please provide the arguments for the command")
	}

	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	queryParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    uuid.NullUUID{UUID: currentUser.ID, Valid: true},
	}

	feed, err := s.db.CreateFeed(context.Background(), queryParams)
	if err != nil {
		return fmt.Errorf("cannot create feed -%v", err)
	}

	fmt.Println("Feed created")
	printFeed(feed)
	// Follow every feed which the current user creates
	handlerFollow(s, command{Args: []string{feedURL}}, currentUser)

	return nil
}
