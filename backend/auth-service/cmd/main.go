package main

import (
	"log"

	"smart-split-save/common/config"
	"smart-split-save/common/logger"
	"smart-split-save/internal/handlers"
	"smart-split-save/internal/storage"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env (for development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it — using environment variables")
	}

	// Load configuration
	cfg := config.Load()

	// Initialize logger
	lg := logger.New()
	defer lg.Sync()

	// Connect to PostgreSQL
	db, err := storage.NewPostgres(cfg.DBUrl)
	if err != nil {
		lg.Sugar().Fatalf("failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Run migrations
	if err := storage.Migrate(db); err != nil {
		lg.Sugar().Fatalf("failed to run migrations: %v", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Create handler
	h := handlers.NewHandler(db, cfg)

	// Register API routes
	api := router.Group("/api/v1")
	{
		api.POST("/signup", h.Signup)
		api.POST("/login", h.Login)
		api.GET("/health", handlers.Health)
	}

	// Start server
	addr := ":" + cfg.Port
	lg.Sugar().Infow("starting auth-service", "addr", addr)
	if err := router.Run(addr); err != nil {
		lg.Sugar().Fatal(err)
	}
}
