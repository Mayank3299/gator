package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mayank3299/BlogAggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, currentUser database.User) error {
	if len(cmd.Args) < 1 {
		return errors.New("please provide the arguments for the command")
	}

	feedUrl := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("feed not found for url - %v", err)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: currentUser.ID, FeedID: feed.ID})
	if err != nil {
		return fmt.Errorf("feed follow not deleted - %v", err)
	}
	fmt.Println("Feed unfollowed")
	return nil
}
