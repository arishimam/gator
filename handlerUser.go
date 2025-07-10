package main

import (
	"context"
	"fmt"
	"time"

	"github.com/arishimam/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("argument slice is empty")
	}

	user := cmd.args[0]
	fmt.Println("trying to login:")
	fmt.Println(user)

	i, err := s.db.GetUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("failed to get user from db: %w", err)
	}

	fmt.Println(i)
	fmt.Println(err)

	err = s.cfg.SetUser(user)
	if err != nil {
		return fmt.Errorf("Couldn't set current user")
	}

	fmt.Println("user has been set to ", user)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("argument slice is empty")
	}
	if len(cmd.args) < 1 {
		return fmt.Errorf("no name was passed in args!")
	}

	userName := cmd.args[0]
	fmt.Println(userName)

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName,
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}

	s.cfg.SetUser(userName)
	fmt.Println("user has been created and set to: ", userName)
	fmt.Println("user: ", user)

	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Println("* " + user.Name + " (current)")
		} else {
			fmt.Println("* " + user.Name)
		}
	}

	return nil

}
