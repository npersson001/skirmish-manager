package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/npersson001/skirmish-manager/internal/models"
	"github.com/npersson001/skirmish-manager/internal/services"
)

type MatchHandler struct {
	service services.MatchService
}

func NewMatchHandler(
	service services.MatchService,
) *MatchHandler {
	return &MatchHandler{
		service: service,
	}
}

type CreateMatchRequest struct {
	WinnerPlayerID int64     `json:"winner_player_id"`
	StartedAt      time.Time `json:"started_at"`
	EndedAt        time.Time `json:"ended_at"`
	PlayerIDs      []int64   `json:"player_ids"`
}

// CreateMatch Creates a new match.
//
// @Summary Create a match
// @Description Creates a new match and records all participating players.
// @Tags matches
// @Accept json
// @Produce json
// @Param match body CreateMatchRequest true "Match to create"
// @Success 201 {object} models.Match
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /matches [post]
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

	match := models.Match{
		WinnerPlayerID: req.WinnerPlayerID,
		StartedAt:      req.StartedAt,
		EndedAt:        req.EndedAt,
	}

	err := h.service.CreateMatch(
		r.Context(),
		&match,
		req.PlayerIDs,
	)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError, // TODO figure out how to get different http status errors for bad request (fails validation in service layer vs failure in repository layer / db issue
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(match); err != nil {
		http.Error(
			w,
			"failed to write response",
			http.StatusInternalServerError,
		)
	}
}

// ListMatches Lists all matches.
//
// @Summary List matches
// @Description Returns all recorded matches.
// @Tags matches
// @Produce json
// @Success 200 {array} models.Match
// @Failure 500 {string} string
// @Router /matches [get]
func (h *MatchHandler) ListMatches(
	w http.ResponseWriter,
	r *http.Request,
) {
	var err error
	var matches []models.Match

	if matches, err = h.service.ListMatches(
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

// GetMatch Gets a match by ID.
//
// @Summary Get a match
// @Description Returns a match by its ID.
// @Tags matches
// @Produce json
// @Param id path int true "Match ID"
// @Success 200 {object} models.Match
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /matches/{id} [get]
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
	} else if match, err = h.service.GetMatch(
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

// DeleteMatch deletes a match.
//
// @Summary Delete a match
// @Description Deletes a match by ID. Intended primarily for administrative use.
// @Tags matches
// @Produce json
// @Param id path int true "Match ID"
// @Success 204
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /matches/{id} [delete]
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
	} else if err = h.service.DeleteMatch(
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
