package dto

import (
	"time"

	"github.com/paulmach/orb"
)

type Party struct {
	ID              string
	UserId          string
	Title           string
	Description     string
	IsPublic        bool
	MaxParticipants int32
	Location        orb.Point
	StreetAddress   string
	PostalCode      string
	State           string
	Country         string
	EntryDate       time.Time
}
