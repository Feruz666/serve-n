package server

import (
	"log"
	"sync"

	"github.com/Feruz666/serve-n/gateway/config"
)

func RunServer(wg *sync.WaitGroup, cfg *config.Config) {
	if cfg == nil {
		log.Fatal("RunServer error", "error", "config is nil")
		return
	}
	defer wg.Done()

	r := NewRouter(cfg)
	err := r.Run(cfg.ServerHost)
	log.Fatal("can not serve", "error", err)
}
