package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	// check if arg's slice is empty
	if len(cmd.Args) == 0 {
		return errors.New("Command's arg's slice is empty")
	}

	// set the user with the given username
	err := s.Config.SetUser(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Println("User has been set")
	return nil
}
