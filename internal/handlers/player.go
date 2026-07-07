package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/npersson001/skirmish-manager/internal/models"
	"github.com/npersson001/skirmish-manager/internal/repository"
)

type PlayerHandler struct {
	repo *repository.PlayerRepository
}

func NewPlayerHandler(
	repo *repository.PlayerRepository,
) *PlayerHandler {
	return &PlayerHandler{
		repo: repo,
	}
}

type CreatePlayerRequest struct {
	Username    string `json:"username"`
	Description string `json:"description"`
}

type UpdatePlayerRequest struct {
	Username    string `json:"username"`
	Description string `json:"description"`
}

func (h *PlayerHandler) CreatePlayer(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req CreatePlayerRequest
	var err error

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	if req.Username == "" {
		http.Error(
			w,
			"username is required",
			http.StatusBadRequest,
		)
		return
	}

	player := models.Player{
		Username:    req.Username,
		Description: req.Description,
	}

	if err = h.repo.Create(
		r.Context(),
		&player,
	); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(player); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
}

func (h *PlayerHandler) ListPlayers(
	w http.ResponseWriter,
	r *http.Request,
) {
	var err error
	var players []models.Player

	if players, err = h.repo.List(
		r.Context(),
	); err != nil {
		fmt.Print(err.Error())
		http.Error(
			w,
			"failed to fetch players",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	if err = json.NewEncoder(w).Encode(players); err != nil {
		fmt.Print(err.Error())
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
}

func (h *PlayerHandler) GetPlayer(
	w http.ResponseWriter,
	r *http.Request,
) {
	idString := chi.URLParam(r, "id")
	var id int64
	var err error
	var player *models.Player

	if id, err = strconv.ParseInt(
		idString,
		10,
		64,
	); err != nil {
		http.Error(
			w,
			"invalid player id",
			http.StatusBadRequest,
		)
		return
	} else if player, err = h.repo.Get(
		r.Context(),
		id,
	); err != nil {
		http.Error(
			w,
			"failed to fetch players",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	if err = json.NewEncoder(w).Encode(player); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
}

func (h *PlayerHandler) UpdatePlayer(
	w http.ResponseWriter,
	r *http.Request,
) {
	idString := chi.URLParam(r, "id")
	var id int64
	var err error
	var player *models.Player

	if id, err = strconv.ParseInt(
		idString,
		10,
		64,
	); err != nil {
		http.Error(
			w,
			"invalid player id",
			http.StatusBadRequest,
		)
		return
	}

	var req UpdatePlayerRequest

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	player = &models.Player{
		ID:          id,
		Username:    req.Username,
		Description: req.Description,
	}

	if err = h.repo.Update(
		r.Context(),
		player,
	); err != nil {
		http.Error(
			w,
			"failed to update player",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	if err = json.NewEncoder(w).Encode(player); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
}

func (h *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	var id int64
	var err error

	if id, err = strconv.ParseInt(
		idString,
		10,
		64,
	); err != nil {
		http.Error(
			w,
			"invalid player id",
			http.StatusBadRequest,
		)
		return
	} else if err = h.repo.Delete(
		r.Context(),
		id,
	); err != nil {
		http.Error(
			w,
			"failed to delete player",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
