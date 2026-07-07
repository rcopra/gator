package main

import (
	"context"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	s.db.DeleteAllUsers(ctx)
	return nil
}
