package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/strattasphere/video-archiver/config"
	"github.com/strattasphere/video-archiver/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	c := config.Load()
	s := server.NewServer(c)
	if err := s.Start(ctx); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
	os.Exit(0)
}
