package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/npersson001/skirmish-manager/internal/models"
	"github.com/npersson001/skirmish-manager/internal/repository"
)

type MatchHandler struct {
	repo *repository.MatchRepository
}

func NewMatchHandler(
	repo *repository.MatchRepository,
) *MatchHandler {
	return &MatchHandler{
		repo: repo,
	}
}

type CreateMatchRequest struct {
	WinnerPlayerID int64     `json:"winner_player_id"`
	StartedAt      time.Time `json:"started_at"`
	EndedAt        time.Time `json:"ended_at"`
	PlayerIDs      []int64   `json:"player_ids"`
}

func (h *MatchHandler) CreateMatch(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req CreateMatchRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	if len(req.PlayerIDs) < 2 {
		http.Error(
			w,
			"a match must contain at least two players",
			http.StatusBadRequest,
		)
		return
	}

	if req.EndedAt.Before(req.StartedAt) {
		http.Error(
			w,
			"ended_at must be after started_at",
			http.StatusBadRequest,
		)
		return
	}

	foundWinner := false
	for _, playerID := range req.PlayerIDs {
		if playerID == req.WinnerPlayerID {
			foundWinner = true
			break
		}
	}

	if !foundWinner {
		http.Error(
			w,
			"winner_player_id must be one of the participating players",
			http.StatusBadRequest,
		)
		return
	}

	match := models.Match{
		WinnerPlayerID: req.WinnerPlayerID,
		StartedAt:      req.StartedAt,
		EndedAt:        req.EndedAt,
	}

	createdMatch, err := h.repo.Create(
		r.Context(),
		&match,
		req.PlayerIDs,
	)

	if err != nil {
		http.Error(
			w,
			"failed to create match",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createdMatch); err != nil {
		http.Error(
			w,
			"failed to write response",
			http.StatusInternalServerError,
		)
	}
}

func (h *MatchHandler) ListMatches(
	w http.ResponseWriter,
	r *http.Request,
) {
	var err error
	var matches []models.Match

	if matches, err = h.repo.List(
		r.Context(),
	); err != nil {
		fmt.Print(err.Error())
		http.Error(
			w,
			"failed to fetch matches",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	if err = json.NewEncoder(w).Encode(matches); err != nil {
		fmt.Print(err.Error())
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
}

func (h *MatchHandler) GetMatch(
	w http.ResponseWriter,
	r *http.Request,
) {
	idString := chi.URLParam(r, "id")
	var id int64
	var err error
	var match *models.Match

	if id, err = strconv.ParseInt(
		idString,
		10,
		64,
	); err != nil {
		http.Error(
			w,
			"invalid match id",
			http.StatusBadRequest,
		)
		return
	} else if match, err = h.repo.Get(
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

	if err = json.NewEncoder(w).Encode(match); err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
	}
}

// TODO make delete an admin operation, treat matches as immutable historical events
func (h *MatchHandler) DeleteMatch(
	w http.ResponseWriter,
	r *http.Request,
) {
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
			"invalid match id",
			http.StatusBadRequest,
		)
		return
	} else if err = h.repo.Delete(
		r.Context(),
		id,
	); err != nil {
		http.Error(
			w,
			"failed to delete match",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
