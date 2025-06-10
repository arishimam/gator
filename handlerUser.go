package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("argument slice is empty")
	}

	user := cmd.args[0]

	err := s.cfg.SetUser(user)
	if err != nil {
		return fmt.Errorf("Couldn't set current user")
	}

	fmt.Println("user has been set to ", user)
	return nil
}
