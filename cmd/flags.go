package main

import (
	"github.com/urfave/cli/v2"
)

var (
	db   string
	port string

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "dbURL",
			EnvVars:     []string{"DB_URL"},
			Destination: &db,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "port",
			EnvVars:     []string{"PORT"},
			Destination: &port,
			Required:    true,
		},
	}
)
