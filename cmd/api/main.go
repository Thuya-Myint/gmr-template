package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Thuya-Myint/gmr-template/internal/cache"
	"github.com/Thuya-Myint/gmr-template/internal/config"
	"github.com/Thuya-Myint/gmr-template/internal/db"
	"github.com/Thuya-Myint/gmr-template/internal/httpserver"
)

func main() {
	cfg := config.Load()

	mongoClient, err := db.ConnectMongo(cfg.MongoURI)
	if err != nil {
		log.Fatalf("MongoDb connection failed: %v", err)
	}
	log.Println("MongoDB connected")

	redisClient, err := cache.ConnectRedis(
		cfg.RedisAddr,
		cfg.RedisPassword,
		cfg.RedisDB,
	)
	if err != nil {
		log.Fatal("Redis connection failed: ", err)
	}
	log.Println("Redis connected")

	addr := ":" + cfg.Port
	server := httpserver.New(addr)

	log.Printf("Starting %s in %s mode on %s", cfg.AppName, cfg.AppEnv, addr)

	// Start the server in a goroutine
	go func() {
		if err := server.Start(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit // Block until a signal is received

	log.Println("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server: %v", err)
	}

	if err := mongoClient.Disconnect(ctx); err != nil {
		log.Printf("Error disconnecting MongoDb: %v", err)
	}
	if err := redisClient.Close(); err != nil {
		log.Printf("Error closing Redis: %v", err)
	}
	log.Printf("Server stopped")
}
