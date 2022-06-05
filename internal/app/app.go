package app

import (
	"context"
	"errors"
	"fmt"
	"go-api-server-example/internal/services/interrupt"
	"os"
	"os/signal"

	"go-api-server-example/internal/config"
	"go-api-server-example/internal/logging"
	"go-api-server-example/internal/server"
)

func RunApp() {
	var (
		debug bool
		ctx   = context.Background()
	)

	if os.Getenv("DEBUG") == "true" {
		debug = true
	}

	log, err := logging.Configure(debug)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		os.Exit(2)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go interrupt.WaitInterruptSignal(c)

	cfg, err := config.Parse(ctx, os.Args)
	if err != nil {
		if errors.Is(err, config.ErrHelpShown) {
			return
		}

		log.Error(err)
		return
	}

	log.Info("example server starting")

	srv, err := server.NewServer(ctx, cfg.Server.Listen)
	if err != nil {
		log.Errorf("failed to get the server: %s", err)
		return
	}

	srv.Run(ctx)
}
