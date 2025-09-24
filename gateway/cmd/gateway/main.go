package main

import (
	"log"
	"net/http"
	"os"

	"gateway/internal/config"
	"gateway/internal/proxy"
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
