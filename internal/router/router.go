package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/npersson001/skirmish-manager/internal/handlers"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/players", func(r chi.Router) {
		r.Post("/", handlers.CreatePlayer)
		r.Get("/", handlers.ListPlayers)
		r.Get("/{id}", handlers.GetPlayer)
		r.Put("/{id}", handlers.UpdatePlayer)
		r.Delete("/{id}", handlers.DeletePlayer)
	})

	return r
}
