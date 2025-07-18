package main

import (
	"context"
	"fmt"
)

func handlerPrintFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Feeds: ")
	for i := range feeds {
		fmt.Println("Name: ", feeds[i].Name)
		fmt.Println("URL: ", feeds[i].Url)
		username, err := s.db.GetUserFromId(context.Background(), feeds[i].UserID)
		if err != nil {
			return err
		}
		fmt.Println("username: ", username)
		fmt.Println()

	}
	return nil
}
