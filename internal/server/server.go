package server

import (
	"net/http"

	"github.com/Yanisssssse/vidego/internal/api"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router chi.Router
	Config *Config
}

func NewServer(c *Config) *Server {
	return &Server{
		Router: api.NewAppRouter(),
		Config: c,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(s.Config.Host+":"+s.Config.Port, s.Router)
}
