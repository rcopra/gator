package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rcopra/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.Args) > 0 {
		parsed, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			return err
		}
		limit = int32(parsed)
	}
	ctx := context.Background()

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}

	posts, err := s.db.GetPostsForUser(ctx, params)
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Println(post.Title)
	}
	return nil
}
