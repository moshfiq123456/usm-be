package app

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/moshfiq123456/ums-be/internal/config"
	"gofr.dev/pkg/gofr"
)

type Server struct {
	app *gofr.App
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	os.Setenv("HTTP_PORT", strconv.Itoa(cfg.Port))
	os.Setenv("PORT", strconv.Itoa(cfg.Port))
	
	app := gofr.New()

	return &Server{
		app: app,
		cfg: cfg,
	}
}

func (s *Server) Start() {
	s.app.GET("/health", func(ctx *gofr.Context) (interface{}, error) {
		return map[string]string{
			"status": "healthy",
			"port":   fmt.Sprintf("%d", s.cfg.Port),
		}, nil
	})

	log.Printf("Server starting on port %d...\n", s.cfg.Port)
	s.app.Run()
}