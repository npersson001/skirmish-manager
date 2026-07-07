package router

import (
	"github.com/go-chi/chi/v5"

	"github.com/npersson001/skirmish-manager/internal/handlers"
)

func New(ph *handlers.PlayerHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/players", func(r chi.Router) {
		r.Post("/", ph.CreatePlayer)
		r.Get("/", ph.ListPlayers)
		r.Get("/{id}", ph.GetPlayer)
		r.Put("/{id}", ph.UpdatePlayer)
		r.Delete("/{id}", ph.DeletePlayer)
	})

	return r
}
