package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/whotterre/odysseus/src/internal/config"
	"github.com/whotterre/odysseus/src/internal/initializers"
	"github.com/whotterre/odysseus/src/internal/models"
	"github.com/whotterre/odysseus/src/internal/routes"
)

func main() {
	app := gin.Default()

	app.Use(cors.Default()) // TODO: Change me

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Failed to load config because %s", err.Error())
	}
	db, err := initializers.ConnectToDB(cfg.DatabaseURL)
	if err != nil {
		log.Printf("Failed to connect to db because %s", err.Error())
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Printf("Failed to migrate models to db because %s", err.Error())
	}
	routes.SetupRoutes(app, db, cfg)

	server := &http.Server{
		Addr:    ":" + cfg.Server.HTTPPort,
		Handler: app,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("HTTP server listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server failed: %s", err.Error())
		}
	}()

	<-ctx.Done()
	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Shutting down HTTP server")
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP server shutdown failed: %s", err.Error())
	}

	log.Println("HTTP server stopped")
}
