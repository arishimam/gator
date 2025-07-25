package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {

	if len(cmd.args) != 1 {
		return fmt.Errorf("This command takes 1 argument: a duration string like '1s, 1m, 1h'")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenReqs)

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	for ; ; <-ticker.C {
		scrapeFeeds(s)

	}

}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	s.db.MarkFeedFetched(context.Background(), feed.ID)
	updatedFeed, err := s.db.GetFeedWithUrl(context.Background(), feed.Url)

	fmt.Println("Just updated ")

	fmt.Println("Feed: ", updatedFeed.Name)
	// fmt.Println("Article: \n", updatedFeed.)

	return nil

}
