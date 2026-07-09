package main

import (
	"context"
	"fmt"

	"github.com/rcopra/gator/internal/database"
)

func handleBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("invalid input")
	}
	ctx := context.Background()
	params, err := GetPostsForUserParams(user.ID, limit: 2)
	posts, err := s.db.GetPostsForUser(ctx, user.ID)
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Println(post.Title)
	}
	return nil
}
