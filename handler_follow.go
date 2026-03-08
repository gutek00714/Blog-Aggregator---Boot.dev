package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gutek00714/Blog-Aggregator---Boot.dev/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: follow <url>")
	}

	url := cmd.Args[0]

	// get feed from url
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	// // get user id
	// user, err := s.db.GetUser(context.Background(), s.Config.CurrentUserName)
	// if err != nil {
	// 	return err
	// }

	follow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %v\n", follow.FeedName)
	fmt.Printf("User: %v\n", follow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	// // get user
	// user, err := s.db.GetUser(context.Background(), s.Config.CurrentUserName)
	// if err != nil {
	// 	return err
	// }

	rows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, row := range rows {
		fmt.Printf("%v\n", row.FeedsName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: unfollow <url>")
	}

	// get the feed by URL
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("feed not found: %w", err)
	}

	// delete the follow
	err = s.db.FeedUnfollow(context.Background(), database.FeedUnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println("Unfollowed successfully")
	return nil
}
