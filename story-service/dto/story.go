package dto

import "github.com/segmentio/ksuid"

type Story struct {
	Id            ksuid.KSUID
	PartyId       string
	UserId        string
	Url           string
	TaggedFriends []string
}
