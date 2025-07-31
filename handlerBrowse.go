package main

import (
	"fmt"
	"strconv"
)

func handlerBrowse(s *state, cmd command) error {

	limit := 2

	if len(cmd.args) >= 1 {
		tmp, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("Argument is not a valid limit. Please use a number!")
		}
		limit = tmp
	}
	fmt.Println(limit)

	return nil

}
