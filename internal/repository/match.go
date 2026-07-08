package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/npersson001/skirmish-manager/internal/models"
)

type MatchRepository struct {
	db *sqlx.DB
}

func NewMatchRepository(db *sqlx.DB) *MatchRepository {
	return &MatchRepository{
		db: db,
	}
}

func (r *MatchRepository) Create(
	ctx context.Context,
	match *models.Match,
	playerIDs []int64,
) (*models.Match, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // if txn fails because of conflict, rollback the changes

	query := `
		INSERT INTO matches (
			winner_player_id,
			started_at,
			ended_at
		)
		VALUES (?, ?, ?)
	`

	result, err := tx.ExecContext(
		ctx,
		query,
		match.WinnerPlayerID,
		match.StartedAt,
		match.EndedAt,
	)

	if err != nil {
		return nil, err
	}

	matchID, err := result.LastInsertId() // TODO does this ALWAYS get the most recent without issue inside txn
	if err != nil {
		return nil, err
	}

	match.ID = matchID

	for _, playerID := range playerIDs {

		_, err := tx.ExecContext(
			ctx,
			`
			INSERT INTO match_players (
				match_id,
				player_id
			)
			VALUES (?, ?)
			`,
			matchID,
			playerID,
		)

		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return match, nil
}

func (r *MatchRepository) List(
	ctx context.Context,
) ([]models.Match, error) {

	var matches []models.Match

	err := r.db.SelectContext(
		ctx,
		&matches,
		`
		SELECT
			id,
			winner_player_id,
			started_at,
			ended_at,
			created_at
		FROM matches
		ORDER BY id
		`,
	)

	if err != nil {
		return nil, err
	}

	return matches, nil
}

func (r *MatchRepository) Get(
	ctx context.Context,
	id int64,
) (*models.Match, error) {

	var match models.Match

	err := r.db.GetContext(
		ctx,
		&match,
		`
		SELECT
			id,
			winner_player_id,
			started_at,
			ended_at,
			created_at
		FROM matches
		WHERE id = ?
		`,
		id,
	)

	if err != nil {
		return nil, err
	}

	return &match, nil
}

func (r *MatchRepository) Update(
	ctx context.Context,
	match *models.Match,
	playerIDs []int64,
) error {

	_, err := r.db.ExecContext(
		ctx,
		`
		UPDATE matches
		SET
			winner_player_id = ?,
			started_at = ?,
			ended_at = ?
		WHERE id = ?
		`,
		match.WinnerPlayerID,
		match.StartedAt,
		match.EndedAt,
		match.ID,
	)

	return err
}

func (r *MatchRepository) Delete(
	ctx context.Context,
	id int64,
) error {

	_, err := r.db.ExecContext(
		ctx,
		`
		DELETE FROM matches
		WHERE id = ?
		`,
		id,
	)

	return err
}

// TODO create table like this:
//FOREIGN KEY (match_id)
//REFERENCES matches(id)
//ON DELETE CASCADE
