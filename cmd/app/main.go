package main

import (
	"context"
	"github.com/Reza-Rayan/internal/config"
	"github.com/Reza-Rayan/internal/db"
	"github.com/Reza-Rayan/internal/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Load Config
	cfg := config.LoadConfig()

	// Init Database
	db.InitDB()

	// Setup Router
	r := routes.SetupRouter()

	// Create HTTP server
	srv := &http.Server{
		Addr:    cfg.Server.Port,
		Handler: r,
	}

	// Run server in goroutine
	go func() {
		log.Printf("🚀 Server started on %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Could not listen on %s: %v\n", cfg.Server.Port, err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("🛑 Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server forced to shutdown: %v", err)
	}

	log.Println("✅ Server exited properly")
}
