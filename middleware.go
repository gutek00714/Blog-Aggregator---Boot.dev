package main

import (
	"context"

	"github.com/gutek00714/Blog-Aggregator---Boot.dev/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	// return a function that matches the standard handler signature
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}

}
