package api

import (
	"github.com/Yanisssssse/vidego/internal/api/handlers"
	"github.com/Yanisssssse/vidego/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func newRouter() chi.Router {
	r := chi.NewRouter()
	return r
}

func NewAppRouter(storage storage.Storage) chi.Router {
	r := newRouter()
	r.Use(middleware.Logger)

	vh := handlers.NewVideoHandlers(storage)

	r.Route("/videos", func(r chi.Router) {
		r.Mount("/", NewVideoRouter(vh))
	})

	return r
}

func NewVideoRouter(h *handlers.VideoHandlers) chi.Router {
	r := newRouter()

	r.Post("/upload", h.Upload)

	return r
}
