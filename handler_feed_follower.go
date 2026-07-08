package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/rcopra/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("invalid input")
	}
	ctx := context.Background()

	url := cmd.Args[0]

	selectedFeed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return err
	}

	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    selectedFeed.ID,
	}

	feed_follows, err := s.db.CreateFeedFollow(ctx, params)
	if err != nil {
		return err
	}

	fmt.Println(feed_follows.FeedName)
	fmt.Println(feed_follows.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("invalid input")
	}
	ctx := context.Background()

	userFeeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return err
	}
	for _, userFeed := range userFeeds {
		fmt.Println(userFeed.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("invalid input")
	}
	ctx := context.Background()

	url := cmd.Args[0]

	selectedFeed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: selectedFeed.ID,
	}

	err = s.db.DeleteFeedFollow(ctx, params)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully deleted %v", selectedFeed.Name)
	return nil
}
