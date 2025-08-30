package main

import (
	"context"
	"fmt"

	"github.com/Mayank3299/BlogAggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, currentUser database.User) error {
	following, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("feed following not found for user- %v", err)
	}

	fmt.Printf("-- %v Feeds --\n", s.configFile.CurrentUser)
	for _, follow := range following {
		fmt.Printf(" * %v\n", follow.FeedName)
	}
	return nil
}
