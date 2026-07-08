package main

import (
	"context"
	"fmt"
	"time"
)

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	markedFeed, err := s.db.MarkFeedFetched(ctx, nextFeed.ID)
	if err != nil {
		return err
	}

	fetchedFeed, err := fetchFeed(ctx, markedFeed.Url)
	if err != nil {
		return err
	}

	fmt.Printf("Fetching feed: %s\n", markedFeed.Name)
	for _, item := range fetchedFeed.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("invalid input missing command")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		fmt.Println("Making request...")
		scrapeFeeds(s)
		fmt.Print("Success")
	}
}
