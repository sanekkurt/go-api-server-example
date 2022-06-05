package common

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"go-api-server-example/internal/server/common/middleware"

	chimiddleware "github.com/go-chi/chi/middleware"
)

const (
	maxRequestsLimit = 100
)

var (
	requestTimeout = 60 * time.Second
)

func ConfigHandlers(r chi.Router) {
	// Use middleware
	r.Use(chimiddleware.Throttle(maxRequestsLimit))
	r.Use(chimiddleware.Timeout(requestTimeout))

	r.Use(middleware.HealthcheckMiddleware, middleware.LoggerMiddleware)

	r.Group(func(r chi.Router) {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))
	})
}
