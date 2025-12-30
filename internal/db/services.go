package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	ID        int
	Name      string
	URL       string
	Status    string
	LastCheck time.Time
}

// CreateService adds a new URL to monitor 
func CreateService(pool *pgxpool.Pool, name, url string) error {
	query := `INSERT INTO services (name, url) VALUES ($1, $2)`
	_, err := pool.Exec(context.Background(), query, name, url)
	return err
}

// GetAllServices returns everything 
func GetAllServices(pool *pgxpool.Pool) ([]Service, error) {
	query := `SELECT id, name, url, status, last_check FROM services`
	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []Service
	for rows.Next() {
		var s Service
		err := rows.Scan(&s.ID, &s.Name, &s.URL, &s.Status, &s.LastCheck)
		if err != nil {
			return nil, err
		}
		services = append(services, s)
	}
	return services, nil
}

// UpdateServiceStatus updates the UP/DOWN status 
func UpdateServiceStatus(pool *pgxpool.Pool, id int, status string) error {
	query := `UPDATE services SET status = $1, last_check = $2 WHERE id = $3`
	_, err := pool.Exec(context.Background(), query, status, time.Now(), id)
	return err
}

func DeleteService(pool *pgxpool.Pool, id string) error {
	_, err := pool.Exec(context.Background(), "DELETE FROM services WHERE id = $1", id)
	return err
}