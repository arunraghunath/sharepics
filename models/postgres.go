package models

import (
	"database/sql"
	"fmt"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
	SSLMode  string
}

func DefaultPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "test",
		Password: "test",
		DBname:   "test",
		SSLMode:  "disable",
	}
}
func (pg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pg.Host, pg.Port, pg.User, pg.Password, pg.DBname, pg.SSLMode)
}

func Open(pg PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", pg.String())
	if err != nil {
		panic(err)
	}
	return db, nil
}
