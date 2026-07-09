package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"

	"github.com/lib/pq"
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
		publishedAt, err := time.Parse(time.RFC822, item.PubDate)
		if err != nil {
			publishedAt, err = time.Parse(time.RFC1123, item.PubDate)
			if err != nil {
				log.Println(err)
				continue
			}
		}

		params := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  err == nil,
			},
			PublishedAt: sql.NullTime{
				Time:  publishedAt,
				Valid: err == nil,
			},
			FeedID: markedFeed.ID,
		}
		_, err = s.db.CreatePost(ctx, params)
		if pqErr, ok := err.(*pq.Error); ok {
			// pqErr is now usable as *pq.Error, ok tells you whether the assertion succeeded
			if pqErr.Code == "23505" {
				continue
			}
			log.Println(pqErr)
		} else if err != nil {
			log.Println(err)
		}

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
