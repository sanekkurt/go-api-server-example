package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go-api-server-example/internal/logging"
)

var (
	shutdownTimeout = 5 * time.Second
)

type Server struct {
	httpServer *http.Server
	listen     int
}

func NewServer(ctx context.Context, listenPort int) (*Server, error) {
	var (
		srv = &Server{
			listen: listenPort,
		}
	)

	srv.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%d", srv.listen),
		Handler:           srv.routes(),
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	return srv, nil
}

func (s *Server) Run(ctx context.Context) {
	var (
		log = logging.GetLoggerFromContext(ctx)
	)
	log.Infof("server started and listen on '%d' port", s.listen)

	go func() {
		<-ctx.Done()

		log.Infof("shutdown initiated")

		shutdownCtx, done := context.WithTimeout(context.Background(), shutdownTimeout)

		defer done()

		if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
			log.Errorf("http shutdown error: %s", err)
			return
		}

		log.Debugf("shutdown completed")
	}()

	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Errorf("failed to serve: %s", err)
		return
	}

	log.Info("stop http server")
}
