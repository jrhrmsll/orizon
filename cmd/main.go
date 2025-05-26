package main

import (
	"context"
	"log"
	"os"
	"syscall"

	"github.com/oklog/run"

	"github.com/jrhrmsll/orizon/pkg/core"
	"github.com/jrhrmsll/orizon/pkg/http"
)

func main() {
	ctx := context.Background()

	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	logger.Println("Starting orizon service")

	// services
	intervalService := core.NewIntervalService()

	var actors run.Group
	{
		signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL}
		actors.Add(run.SignalHandler(ctx, signals...))

		app := http.NewApplication(ctx, logger, intervalService)
		actors.Add(app.Start, app.Stop)
	}

	if err := actors.Run(); err != nil {
		logger.Println(err)
	}
}
