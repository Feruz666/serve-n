package main

import (
	"fmt"
	"sync"

	"github.com/Feruz666/serve-n/gateway/config"
	"github.com/Feruz666/serve-n/gateway/server"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	cfg, err := config.InitConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go server.RunServer(&wg, cfg)

	wg.Wait()

	return nil
}
