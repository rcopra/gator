package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/rcopra/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("invalid input")
	}
	ctx := context.Background()

	currentUser, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    currentUser.ID,
	}

	feed, err := s.db.CreateFeed(ctx, feedParams)
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    feed.UserID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(ctx, feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Println(feed.ID)
	fmt.Println(feed.CreatedAt)
	fmt.Println(feed.UpdatedAt)
	fmt.Println(feed.Name)
	fmt.Println(feed.Url)
	fmt.Println(feed.UserID)
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	ctx := context.Background()
	feedSlice, err := s.db.GetFeeds(ctx)
	if err != nil {
		return err
	}
	for _, feed := range feedSlice {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(feed.UserName)
	}

	return nil
}
