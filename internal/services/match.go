package services

import (
	"context"
	"errors"

	"github.com/npersson001/skirmish-manager/internal/models"
	"github.com/npersson001/skirmish-manager/internal/repository"
)

type MatchService interface {
	CreateMatch(ctx context.Context, match *models.Match, playerIds []int64) error
	ListMatches(ctx context.Context) ([]models.Match, error)
	GetMatch(ctx context.Context, id int64) (*models.Match, error)
	DeleteMatch(ctx context.Context, id int64) error
}

type MatchServiceImpl struct {
	repo repository.MatchRepository
}

func NewMatchService(
	repo repository.MatchRepository,
) MatchService {
	return &MatchServiceImpl{
		repo: repo,
	}
}

func (s *MatchServiceImpl) CreateMatch(
	ctx context.Context,
	match *models.Match,
	playerIds []int64,
) error {

	if len(playerIds) < 2 {
		return errors.New("a match must contain at least two players")
	}

	if match.EndedAt.Before(match.StartedAt) {
		return errors.New("ended_at must be after started_at")
	}

	foundWinner := false
	for _, playerID := range playerIds {
		if playerID == match.WinnerPlayerID {
			foundWinner = true
			break
		}
	}

	if !foundWinner {
		return errors.New("winner_player_id must be one of the participating players")
	}

	return s.repo.Create(ctx, match, playerIds)
}

func (s *MatchServiceImpl) ListMatches(
	ctx context.Context,
) ([]models.Match, error) {
	return s.repo.List(ctx)
}

func (s *MatchServiceImpl) GetMatch(
	ctx context.Context,
	id int64,
) (*models.Match, error) {
	return s.repo.Get(ctx, id)
}

func (s *MatchServiceImpl) DeleteMatch(
	ctx context.Context,
	id int64,
) error {
	return s.repo.Delete(ctx, id)
}
