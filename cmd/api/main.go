package main

import (
	"log"
	"net/http"

	"github.com/npersson001/skirmish-manager/internal/handlers"
	"github.com/npersson001/skirmish-manager/internal/repository"
	"github.com/npersson001/skirmish-manager/internal/router"
	"github.com/npersson001/skirmish-manager/internal/services"
)

func main() {
	db, err := repository.NewMySQLConnection()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	playerRepo := repository.NewPlayerRepository(db)
	playerService := services.NewPlayerService(playerRepo)

	playerHandler := handlers.NewPlayerHandler(
		playerService,
	)

	matchRepo := repository.NewMatchRepository(db)
	matchService := services.NewMatchService(matchRepo)

	matchHandler := handlers.NewMatchHandler(
		matchService,
	)

	r := router.New(playerHandler, matchHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
