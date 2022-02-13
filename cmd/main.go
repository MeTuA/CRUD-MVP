package main

import (
	"context"
	"crud/mvp/users"
	"os"
	"sync"

	"github.com/hashicorp/go-hclog"
	"github.com/urfave/cli/v2"
)

func main() {

	log := hclog.New(&hclog.LoggerOptions{
		Name:       "crud/mvp/main.go",
		JSONFormat: true,
		Level:      hclog.Debug,
	})

	app := &cli.App{
		Name:  "crud/mvp",
		Usage: "go run .",
		Flags: flags,
	}

	app.Action = func(c *cli.Context) error {
		connection, err := connectDB(log.Named("DB"))
		if err != nil {
			log.Error("failed to connect DB", "error", err.Error())
			os.Exit(1)
		}

		ctx, cancel := context.WithCancel(context.Background())
		wg := &sync.WaitGroup{}

		usersStore := users.NewStore(connection, log.Named("usersStore"))
		usersService := users.NewService(usersStore, log.Named("usersService"))

		startService(ctx, wg, usersService, log.Named("startService"))
		waitForExit(cancel, wg, log.Named("waitForExit"))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error("failed to run app", "error", err.Error())
		os.Exit(1)
	}
}
