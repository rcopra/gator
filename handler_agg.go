package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	ctx := context.Background()
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(ctx, feedURL)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
