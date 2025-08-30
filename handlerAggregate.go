package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Mayank3299/BlogAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAggregate(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}
	log.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

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
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}
		// fmt.Printf("Found post: %s\n", item.Title)
		queryParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       sql.NullString{String: item.Title, Valid: true},
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: publishedAt,
			FeedID:      updatedFeed.ID,
		}
		_, err := s.db.CreatePost(context.Background(), queryParams)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}

	log.Printf("Feed %s collected, %v posts found", updatedFeed.Name, len(feedData.Channel.Item))
	return nil
}
