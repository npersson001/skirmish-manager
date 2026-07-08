package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/npersson001/skirmish-manager/internal/models"
)

type PlayerRepository interface {
	Create(ctx context.Context, player *models.Player, ) error
	Get(ctx context.Context, id int64) (*models.Player, error)
	List(ctx context.Context, ) ([]models.Player, error)
	Update(ctx context.Context, player *models.Player, ) error
	Delete(ctx context.Context, id int64, ) error
}

type PlayerRepositoryImpl struct {
	db *sqlx.DB
}

func NewPlayerRepository(db *sqlx.DB) PlayerRepository {
	return &PlayerRepositoryImpl{
		db: db,
	}
}

func (r *PlayerRepositoryImpl) Create(
	ctx context.Context,
	player *models.Player,
) error {
	query := `
		INSERT INTO players (
			username,
			description
		)
		VALUES (?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		player.Username,
		player.Description,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	player.ID = id

	return nil
}

func (r *PlayerRepositoryImpl) List(
	ctx context.Context,
) ([]models.Player, error) {

	query := `
		SELECT id, username, description, created_at
		FROM players
		ORDER BY id
	`

	var players []models.Player

	err := r.db.SelectContext(
		ctx,
		&players,
		query,
	)

	if err != nil {
		return nil, err
	}

	return players, nil
}

func (r *PlayerRepositoryImpl) Get(
	ctx context.Context,
	id int64,
) (*models.Player, error) {

	query := `
		SELECT id, username, description, created_at
		FROM players
		WHERE id = ?
	`

	var player models.Player

	err := r.db.GetContext(
		ctx,
		&player,
		query,
		id,
	)

	if err != nil {
		return nil, err
	}

	return &player, nil
}

func (r *PlayerRepositoryImpl) Update(
	ctx context.Context,
	player *models.Player,
) error {

	query := `
		UPDATE players
		SET username = ?, description = ?
		WHERE id = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		player.Username,
		player.Description,
		player.ID,
	)

	return err
}

func (r *PlayerRepositoryImpl) Delete(
	ctx context.Context,
	id int64,
) error {

	query := `
		DELETE FROM players
		WHERE id = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		id,
	)

	return err
}
