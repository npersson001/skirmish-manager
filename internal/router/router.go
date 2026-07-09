package router

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/npersson001/skirmish-manager/docs"

	"github.com/npersson001/skirmish-manager/internal/handlers"
)

func New(ph *handlers.PlayerHandler, mh *handlers.MatchHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/players", func(r chi.Router) {
		r.Post("/", ph.CreatePlayer)
		r.Get("/", ph.ListPlayers)
		r.Get("/{id}", ph.GetPlayer)
		r.Put("/{id}", ph.UpdatePlayer)
		r.Delete("/{id}", ph.DeletePlayer)
	})

	r.Route("/matches", func(r chi.Router) {
		r.Post("/", mh.CreateMatch)
		r.Get("/", mh.ListMatches)
		r.Get("/{id}", mh.GetMatch)
		r.Delete("/{id}", mh.DeleteMatch)
	})

	r.Mount("/swagger", httpSwagger.WrapHandler)

	return r
}
