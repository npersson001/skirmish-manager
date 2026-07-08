package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/npersson001/skirmish-manager/internal/models"
	"github.com/npersson001/skirmish-manager/internal/repository"
)

const MAX_DESCRIPTION_LENGTH = 255

type PlayerService interface {
	CreatePlayer(ctx context.Context, player *models.Player) error
	ListPlayers(ctx context.Context) ([]models.Player, error)
	GetPlayer(ctx context.Context, id int64) (*models.Player, error)
	UpdatePlayer(ctx context.Context, player *models.Player) error
	DeletePlayer(ctx context.Context, id int64) error
}

type PlayerServiceImpl struct {
	repo repository.PlayerRepository
}

func NewPlayerService(
	repo repository.PlayerRepository,
) PlayerService {
	return &PlayerServiceImpl{
		repo: repo,
	}
}

func (s *PlayerServiceImpl) CreatePlayer(
	ctx context.Context,
	player *models.Player,
) error {

	if player.Username == "" {
		return errors.New("username is required")
	}

	if len(player.Description) > MAX_DESCRIPTION_LENGTH {
		return errors.New(fmt.Sprintf("description is too long, must be under %d characters", MAX_DESCRIPTION_LENGTH))
	}

	return s.repo.Create(ctx, player)
}

func (s *PlayerServiceImpl) ListPlayers(
	ctx context.Context,
) ([]models.Player, error) {
	return s.repo.List(ctx)
}

func (s *PlayerServiceImpl) GetPlayer(
	ctx context.Context,
	id int64,
) (*models.Player, error) {
	return s.repo.Get(ctx, id)
}

func (s *PlayerServiceImpl) UpdatePlayer(
	ctx context.Context,
	player *models.Player,
) error {
	if len(player.Description) > MAX_DESCRIPTION_LENGTH {
		return errors.New(fmt.Sprintf("description is too long, must be under %d characters", MAX_DESCRIPTION_LENGTH))
	}

	return s.repo.Update(ctx, player)
}

func (s *PlayerServiceImpl) DeletePlayer(
	ctx context.Context,
	id int64,
) error {
	return s.repo.Delete(ctx, id)
}
