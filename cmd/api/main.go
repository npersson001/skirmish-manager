package main

import (
	"log"
	"net/http"

	"github.com/npersson001/skirmish-manager/internal/router"
)

func main() {
	r := router.New()

	log.Println("Listening on :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
