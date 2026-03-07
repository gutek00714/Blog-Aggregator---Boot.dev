package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gutek00714/Blog-Aggregator---Boot.dev/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	// get the user from database
	user := s.Config.CurrentUserName
	user_db, err := s.db.GetUser(context.Background(), user)
	if err != nil {
		return err
	}

	// check if cmf.args has 2 arguments
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %v <name> <url>", cmd.Name)
	}

	// get the data from args
	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user_db.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %v", err)
	}

	fmt.Println("Feed created successfully:")
	fmt.Println(feed)
	return nil
}
