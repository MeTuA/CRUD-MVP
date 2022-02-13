package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/hashicorp/go-hclog"
	"github.com/labstack/echo/v4"
)

func startService(ctx context.Context, wg *sync.WaitGroup, e *echo.Echo, log hclog.Logger) {

	wg.Add(1)
	go func() {
		log.Info("starting service")

		err := e.Start(":" + port)
		if err != nil && err != http.ErrServerClosed {
			log.Error("failed to start service", "error", err.Error())
			os.Exit(1)
		}
		wg.Done()
	}()

	go func() {
		<-ctx.Done()
		err := e.Shutdown(context.Background())
		if err != nil {
			log.Error("failed to shutdown service", "error", err.Error())
		}
	}()
}

func waitForExit(cancel func(), wg *sync.WaitGroup, log hclog.Logger) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	cancel()
	wg.Wait()
	log.Info("service stopped")
}
