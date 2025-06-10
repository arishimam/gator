package main

import (
	"fmt"
	"os"

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

	if len(os.Args) < 2 {
		fmt.Println("No command-line arguments passed in")
		os.Exit(1)
	}
	args := os.Args[1:]
	if len(args) == 1 {
		fmt.Println("username is required")
		os.Exit(1)
	}

	username := args[1]
	s.cfg.SetUser(username)

}
