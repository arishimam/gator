package main

import (
	"context"
	"fmt"
	"github.com/arishimam/gator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerAddFeed(s *state, cmd command) error {
	userName := s.cfg.CurrentUserName
	fmt.Println(userName)

	if len(cmd.args) < 2 {
		return fmt.Errorf("Not enough args")
	}
	feedName := cmd.args[0]
	url := cmd.args[1]
	fmt.Println(feedName, url)

	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       url,
		UserID:    user.ID,
	}

	s.db.CreateFeed(context.Background(), feedParams)

	fmt.Printf("feed: %+v\n", *feed)

	return nil

}
