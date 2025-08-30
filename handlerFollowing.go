package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.configFile.CurrentUser)
	if err != nil {
		return fmt.Errorf("user not found - %v", err)
	}

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
