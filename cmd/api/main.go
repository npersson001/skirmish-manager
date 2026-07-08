package main

import (
	"log"
	"net/http"

	"github.com/npersson001/skirmish-manager/internal/handlers"
	"github.com/npersson001/skirmish-manager/internal/repository"
	"github.com/npersson001/skirmish-manager/internal/router"
)

func main() {
	db, err := repository.NewMySQLConnection()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	playerRepo := repository.NewPlayerRepository(db)

	playerHandler := handlers.NewPlayerHandler(
		playerRepo,
	)

	matchRepo := repository.NewMatchRepository(db)

	matchHandler := handlers.NewMatchHandler(
		matchRepo,
	)

	r := router.New(playerHandler, matchHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
