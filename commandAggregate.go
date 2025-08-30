package main

import (
	"context"
	"fmt"
)

func handlerAggregate(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	rssFeed, err := fetch(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	fmt.Printf("RSS Feed - %+v", rssFeed)
	return nil
}
