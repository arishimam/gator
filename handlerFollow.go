package main

import (
	"context"
	"fmt"
	"time"

	"github.com/arishimam/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("Incorrect number of arguments!")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedWithUrl(context.Background(), url)
	if err != nil {
		return err
	}

	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	newFeedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return err
	}

	fmt.Println("NewFeedFollow name: ", newFeedFollow.FeedName)
	fmt.Println("username: ", newFeedFollow.UserName)

	return nil
}
func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		fmt.Errorf("This command only takes 1 argument")
	}

	url := cmd.args[0]
	feed, err := s.db.GetFeedWithUrl(context.Background(), url)
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}

	return nil

}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("No arguments allowed for this command!")
	}

	username := s.cfg.CurrentUserName

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("%v's Feeds:\n", username)

	for _, ff := range feedFollows {
		fmt.Println(ff.FeedName)
	}
	return nil

}
