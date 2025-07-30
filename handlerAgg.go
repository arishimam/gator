package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/arishimam/gator/internal/database"
	"github.com/google/uuid"
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
	ctx := context.Background()

	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get next feed: %w\n", err)
	}

	if err := s.db.MarkFeedFetched(ctx, feed.ID); err != nil {
		return fmt.Errorf("failed to mark feed fetched: %w\n", err)
	}

	// updatedFeed, err := s.db.GetFeedWithUrl(ctx, feed.Url)

	fetchedFeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("error occurred when fetching feed: %w\n", err)
	}

	//fmt.Println("UPDATED:")
	//fmt.Println(fetchedFeed)

	for _, item := range fetchedFeed.Channel.Item {
		pubDate, err := parsePubDate(item.PubDate)
		if err != nil {
			fmt.Printf("Skipping item due to bad pubDate: %q (%v)\n", item.PubDate, err)
			continue
		}

		postParams := database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     fetchedFeed.Channel.Title,
			Url:       feed.Url,
			Description: sql.NullString{
				String: fetchedFeed.Channel.Description,
				Valid:  fetchedFeed.Channel.Description != ""},
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		}

		_, err = s.db.CreatePost(ctx, postParams)
		if err != nil {
			fmt.Printf("Duplicated or failed insert for URL %s: %v\n", item.Link, err)
		}
	}

	fmt.Println("Scrape complete")

	return nil

}

func parsePubDate(s string) (time.Time, error) {
	formats := []string{
		time.RFC1123,
		time.RFC1123Z,
		time.RFC822Z,
		time.RFC822,
		time.RFC3339,
		time.RFC850,
	}

	for _, f := range formats {
		if t, err := time.Parse(f, s); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unrecognized publish date format: %s\n", s)
}
