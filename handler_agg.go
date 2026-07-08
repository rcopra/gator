package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/rcopra/gator/internal/database"
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
		params := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: sql.NullTime{
				Time:  item.PubDate,
				Valid: true,
			},
			FeedID: markedFeed.ID,
		}

		post, err := s.db.CreatePost(ctx, params)
		if err != nil {
			return err
		}
		fmt.Printf("Successfully saved post: %s", post.Title)

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

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		fmt.Println("Making request...")
		if err := scrapeFeeds(s); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Success")
	}
}
