package models

import "time"

type Player struct {
	ID          int64     `db:"id" json:"id"`
	Username    string    `db:"username" json:"username"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
