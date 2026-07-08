package models

type MatchPlayer struct {
	MatchID  int64 `db:"match_id" json:"match_id"`
	PlayerID int64 `db:"player_id" json:"player_id"`
}
