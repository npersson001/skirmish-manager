package models

import "time"

type Match struct {
	ID             int64     `db:"id" json:"id"`
	WinnerPlayerID int64     `db:"winner_player_id" json:"winner_player_id"`
	StartedAt      time.Time `db:"started_at" json:"started_at"`
	EndedAt        time.Time `db:"ended_at" json:"ended_at"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}
