package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/gutek00714/Blog-Aggregator---Boot.dev/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	// check if arg's slice is empty
	if len(cmd.Args) == 0 {
		return errors.New("Command's arg's slice is empty")
	}

	// check if the user exists in the database
	_, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("user not found: %s", cmd.Args[0])
	}

	// set the user with the given username
	err = s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("a name is required")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %v", err)
	}

	// set the user in the config
	err = s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("User was created: ")
	fmt.Println(user)
	return nil
}
