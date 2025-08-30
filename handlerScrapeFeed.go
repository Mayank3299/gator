package main

import (
	"context"
	"fmt"
	"log"
)

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("feed not fetched - %v", err)
	}

	updatedFeed, err := s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("couldn't mark feed %s fetched: %v", nextFeed.Name, err)
	}

	feedData, err := fetch(context.Background(), updatedFeed.Url)
	if err != nil {
		return fmt.Errorf("couldn't collect feed: %v", err)
	}

	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}

	log.Printf("Feed %s collected, %v posts found", updatedFeed.Name, len(feedData.Channel.Item))
	return nil
}
