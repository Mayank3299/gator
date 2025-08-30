package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Mayank3299/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerBrowse(s *state, cmd command, currentUser database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit %v", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: uuid.NullUUID{UUID: currentUser.ID, Valid: true},
		Limit:  int32(limit),
	})

	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(posts), currentUser.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %v ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}
