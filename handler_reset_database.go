package main

import (
	"context"
	"fmt"
)

func handlerResetDatabase(s *state, cmd command) error {
	err := s.db.ResetDB(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset database: %v", err)
	}

	fmt.Println("Database reset successful")
	return nil
}
