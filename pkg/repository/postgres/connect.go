package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connect establish connection to a DB with a given configuration
func Connect(cfg Config) (*sql.DB, error) {
	postgres, err := sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	))
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := postgres.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return postgres, nil
}
