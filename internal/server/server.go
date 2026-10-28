package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/strattasphere/video-archiver/config"
	"github.com/strattasphere/video-archiver/internal/server/router"
)

type Server struct {
	srv  *http.Server
	conf config.Config
}

func NewServer(c config.Config) *Server {
	r := router.New()

	return &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", c.Host, c.Port),
			Handler: r,
		},
		conf: c,
	}
}

func (s *Server) Start(ctx context.Context) error {
	serveErr := make(chan error, 1)

	go func() { // sender
		log.Println("starting server...")
		if err := s.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			serveErr <- err
		}
	}()

	select { // receiver
	case err := <-serveErr:
		return err
	case <-ctx.Done():
		log.Println("shutdown signal received")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown incomplete: %v", err)
		log.Printf("forcing shutdown")
		if closeErr := s.srv.Close(); closeErr != nil { // force close connections if graceful shutdown fails
			return fmt.Errorf("forced close failed: %w", closeErr)
		}
		return err
	}

	return nil
}
