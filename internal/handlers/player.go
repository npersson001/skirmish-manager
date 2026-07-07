package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/npersson001/skirmish-manager/internal/models"
)

func CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var player models.Player

	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("CreatePlayer: %+v\n", player)

	w.WriteHeader(http.StatusCreated)
}

func ListPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListPlayers")

	w.WriteHeader(http.StatusOK)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	fmt.Printf("GetPlayer: id=%s\n", id)

	w.WriteHeader(http.StatusOK)
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var player models.Player

	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("UpdatePlayer: id=%s player=%+v\n", id, player)

	w.WriteHeader(http.StatusOK)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	fmt.Printf("DeletePlayer: id=%s\n", id)

	w.WriteHeader(http.StatusNoContent)
}
