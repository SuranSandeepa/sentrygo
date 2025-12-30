package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/SuranSandeepa/sentrygo/internal/db"
	"github.com/SuranSandeepa/sentrygo/internal/handlers"
	"github.com/SuranSandeepa/sentrygo/internal/monitor" // 1. MAKE SURE THIS IS HERE
)

func main() {
	pool, err := db.Connect()
	if err != nil {
		log.Fatalf("Critical: Could not connect to database: %v", err)
	}
	defer pool.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	schema := `
	CREATE TABLE IF NOT EXISTS services (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		url TEXT NOT NULL UNIQUE,
		status TEXT DEFAULT 'PENDING',
		last_check TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = pool.Exec(ctx, schema)
	if err != nil {
		log.Fatalf("Critical: Could not create database table: %v", err)
	}
	fmt.Println("✅ Database is ready.")

	// --- NEW PART STARTS HERE ---
	// 2. Start the Background Worker
	// We use 'go' so it runs in the background while the rest of the code continues
	go monitor.StartWorker(pool, 30*time.Second) 
	// --- NEW PART ENDS HERE ---

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", handlers.Dashboard(pool))
	r.Post("/add", handlers.AddService(pool))
	r.Delete("/delete/{id}", handlers.DeleteService(pool))

	port := ":8080"
	fmt.Printf("🚀 SentryGo Dashboard active at http://localhost%s\n", port)
	
	server := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Critical: Server failed: %v", err)
	}
}