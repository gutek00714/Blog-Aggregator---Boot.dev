package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gutek00714/Blog-Aggregator---Boot.dev/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 0 {
		if l, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = l
		}
	}

	posts, err := s.db.GetUserPosts(context.Background(), database.GetUserPostsParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Found %v posts for user %v:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s -- %s\n", post.PublishedAt.Time.Format("Jan 2"), post.Title)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Printf("Description: %s\n", post.Description.String)
	}

	return nil
}
