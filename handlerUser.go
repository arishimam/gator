package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("argument slice is empty")
	}

	user := cmd.args[0]
	err := s.cfg.SetUser(user)
	if err != nil {
		return err
	}
	fmt.Println("user has been set to ", user)

	return nil
}
