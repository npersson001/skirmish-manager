package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/npersson001/skirmish-manager/internal/models"
)

type PlayerRepository struct {
	db *sqlx.DB
}

func NewPlayerRepository(db *sqlx.DB) *PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}

func (r *PlayerRepository) Create(
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

func (r *PlayerRepository) List(
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

func (r *PlayerRepository) Get(
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

func (r *PlayerRepository) Update(
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

func (r *PlayerRepository) Delete(
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
