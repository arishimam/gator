package main

import (
	"fmt"
	"github.com/arishimam/gator/internal/config"
)

type state struct {
	cfg *Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		return
	}

	fmt.Println(cfg)
	cfg.SetUser("arish")
	cfg, _ = config.Read()
	fmt.Println(cfg)
}
