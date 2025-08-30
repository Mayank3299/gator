package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Mayank3299/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return errors.New("please provide the arguments for the command")
	}

	feedUrl := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("feed not found for url - %v", err)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.configFile.CurrentUser)
	if err != nil {
		return fmt.Errorf("user not found - %v", err)
	}

	queryParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    feed.ID,
		UserID:    currentUser.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), queryParams)
	if err != nil {
		return fmt.Errorf("cannot create feed follow - %v", err)
	}

	fmt.Println("FeedFollow created")
	printFeedFollow(feedFollow)
	return nil
}
