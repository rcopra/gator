package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"

	"github.com/rcopra/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("invalid input missing command")
	}
	userName := cmd.Args[0]

	ctx := context.Background()

	_, err := s.db.GetUser(ctx, userName)
	if err != nil {
		return err
	}

	if err := s.cfg.SetUser(userName); err != nil {
		return err
	}
	fmt.Printf("Username successfully set to %s\n", userName)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("invalid input missing command")
	}

	userName := cmd.Args[0]

	ctx := context.Background()

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}

	_, err := s.db.CreateUser(ctx, params)
	if err != nil {
		return err
	}
	if err := s.cfg.SetUser(userName); err != nil {
		return err
	}

	fmt.Printf("Username: %v, successfully added to list of users\n", userName)

	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	ctx := context.Background()
	userSlice, err := s.db.GetUsers(ctx)
	if err != nil {
		return err
	}

	for _, user := range userSlice {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Println(user.Name + " (current)")
		} else {
			fmt.Println(user.Name)
		}
	}

	return nil
}
