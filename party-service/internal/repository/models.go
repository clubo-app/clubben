// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package repository

import (
	"database/sql"
	"time"

	"github.com/paulmach/orb"
)

type Party struct {
	ID                string
	UserID            string
	Title             string
	Description       sql.NullString
	IsPublic          bool
	MusicGenre        sql.NullString
	Location          orb.Point
	StreetAddress     sql.NullString
	PostalCode        sql.NullString
	State             sql.NullString
	Country           sql.NullString
	EntryDate         time.Time
	MaxParticipants   int32
	ParticipantsCount int32
	FavoriteCount     int32
}