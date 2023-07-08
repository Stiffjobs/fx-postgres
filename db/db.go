package db

import (
	"fmt"
	"fx-postgres/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

var DatabaseModule = fx.Provide(NewDatabase)

func NewDatabase(config config.EnvVars) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", config.Username, config.Password, config.Host, config.Port, config.DbName)
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
