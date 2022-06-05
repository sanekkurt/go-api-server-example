package server

import (
	"github.com/go-chi/chi"
	"go-api-server-example/internal/server/common"
)

func (s *Server) routes() chi.Router {
	r := chi.NewRouter()

	common.ConfigHandlers(r)

	return r
}
