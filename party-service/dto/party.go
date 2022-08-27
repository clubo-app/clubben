package dto

import (
	"time"

	"github.com/paulmach/orb"
)

type Party struct {
	ID              string
	UserId          string
	Title           string
	IsPublic        bool
	MaxParticipants uint32
	Location        orb.Point
	StreetAddress   string
	PostalCode      string
	State           string
	Country         string
	StartDate       time.Time
	EndDate         time.Time
}
