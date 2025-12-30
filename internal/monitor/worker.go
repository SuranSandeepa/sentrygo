package monitor

import (
	"fmt"
	"time"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/SuranSandeepa/sentrygo/internal/db"
)

// StartWorker runs the background loop
func StartWorker(pool *pgxpool.Pool, interval time.Duration) {
	ticker := time.NewTicker(interval)
	fmt.Println("🛰️  Background Monitor Started...")

	// Run once immediately on start
	checkAll(pool)

	for range ticker.C {
		checkAll(pool)
	}
}

func checkAll(pool *pgxpool.Pool) {
	services, err := db.GetAllServices(pool)
	if err != nil {
		fmt.Println("Worker Error:", err)
		return
	}

	for _, s := range services {
		// This calls the function from monitor.go
		result := CheckService(s.URL)
		db.UpdateServiceStatus(pool, s.ID, result.Status)
		fmt.Printf("🔄 Checked %s: %s\n", s.Name, result.Status)
	}
}