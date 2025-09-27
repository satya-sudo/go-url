package main

import (
	"log"
	"net/http"

	"github.com/satya-sudo/go-url/gateway/internal/config"
	"github.com/satya-sudo/go-url/gateway/internal/proxy"
)

func main() {
	cfg := config.Load()

	r := proxy.SetupRouter(cfg)

	addr := ":" + cfg.Port
	log.Printf("🚀 Gateway listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
