package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		feedUser, err := s.db.GetUserById(context.Background(), feed.UserID.UUID)
		if err != nil {
			return fmt.Errorf("user not present in db -%v", err)
		}
		fmt.Printf("Name: %v\n", feed.Name)
		fmt.Printf("URL: %v\n", feed.Url)
		fmt.Printf("URL: %v\n", feedUser.Name)
		fmt.Println("---------")
	}

	return nil
}
