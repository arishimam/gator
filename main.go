package main

import (
	"github.com/arishimam/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	s := state{}
	config, err := config.Read()
	s.cfg = &config

	cmds := commands{make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	if err != nil {
		return
	}

}
