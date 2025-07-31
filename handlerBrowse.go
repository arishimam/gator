package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/arishimam/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	limit := 2

	if len(cmd.args) >= 1 {
		tmp, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("Argument is not a valid limit. Please use a number!")
		}
		limit = tmp
	}

	getPostsParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostsForUser(ctx, getPostsParams)
	if err != nil {
		return fmt.Errorf("error getting posts: %w\n", err)
	}
	for _, p := range posts {
		fmt.Println(p.Title)
		fmt.Println(p.Description)
		fmt.Println(p.Url)
		fmt.Println()
	}

	return nil

}
