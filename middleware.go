package main

import (
	"context"

	"github.com/rcopra/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		ctx := context.Background()

		currentUser, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		err = handler(s, cmd, currentUser)
		return err
	}
}
