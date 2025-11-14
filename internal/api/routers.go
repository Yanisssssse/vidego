package api

import (
	"github.com/Yanisssssse/vidego/internal/api/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func newRouter() chi.Router {
	r := chi.NewRouter()
	return r
}

func NewAppRouter() chi.Router {
	r := newRouter()
	r.Use(middleware.Logger)

	r.Route("/videos", func(r chi.Router) {
		r.Mount("/", NewVideoRouter())
	})

	return r
}

func NewVideoRouter() chi.Router {
	r := newRouter()

	r.Post("/upload", handlers.Upload)

	return r
}
