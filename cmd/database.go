package main

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/hashicorp/go-hclog"

	"github.com/jmoiron/sqlx"
)

func connectDB(log hclog.Logger) (*sqlx.DB, error) {

	conn, err := sqlx.Connect("postgres", db)
	if err != nil {
		log.Error("failed to connect DB", "error", err.Error())
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		log.Error("failed to ping DB", "error", err.Error())
		return nil, err
	}

	log.Info("connected to database")

	return conn, nil
}
