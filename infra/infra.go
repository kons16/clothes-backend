package infra

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "test:test@(localhost:3306)/test")
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping: %w", err)
	}

	return db, nil
}
