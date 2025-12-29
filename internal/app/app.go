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
	// Set HTTP port
	os.Setenv("HTTP_PORT", strconv.Itoa(cfg.Port))
	os.Setenv("PORT", strconv.Itoa(cfg.Port))
	
	// Set database configuration for GoFr
	os.Setenv("DB_HOST", cfg.Database.Host)
	os.Setenv("DB_PORT", strconv.Itoa(cfg.Database.Port))
	os.Setenv("DB_USER", cfg.Database.User)
	os.Setenv("DB_PASSWORD", cfg.Database.Password)
	os.Setenv("DB_NAME", cfg.Database.DBName)
	os.Setenv("DB_DIALECT", "postgres")
	
	app := gofr.New()

	return &Server{
		app: app,
		cfg: cfg,
	}
}

func (s *Server) Start() {
	// Health check endpoint
	s.app.GET("/health", func(ctx *gofr.Context) (interface{}, error) {
		return map[string]string{
			"status":   "healthy",
			"port":     fmt.Sprintf("%d", s.cfg.Port),
			"database": s.cfg.Database.DBName,
		}, nil
	})

	// Database health check
	s.app.GET("/health/db", func(ctx *gofr.Context) (interface{}, error) {
		// GoFr automatically provides database connection via ctx.SQL
		_, err := ctx.SQL.Exec("SELECT 1")
		if err != nil {
			return map[string]string{
				"status": "unhealthy",
				"error":  err.Error(),
			}, nil
		}
		
		return map[string]string{
			"status": "healthy",
		}, nil
	})

	log.Printf("Server starting on port %d...\n", s.cfg.Port)
	s.app.Run()
}