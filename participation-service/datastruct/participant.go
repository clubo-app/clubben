package datastruct

import (
	"time"

	"github.com/clubo-app/clubben/protobuf/participation"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Participant struct {
	UserId      string    `db:"user_id"       validate:"required"`
	PartyId     string    `db:"party_id"      validate:"required"`
	Requested   bool      `db:"requested"`
	JoinedAt    time.Time `db:"joined_at"     validate:"required"`
	RequestedAt time.Time `db:"requested_at"  validate:"required"`
}

func (i Participant) ToGRPCPartyParticipant() *participation.PartyParticipant {
	return &participation.PartyParticipant{
		UserId:      i.UserId,
		PartyId:     i.PartyId,
		Requested:   i.Requested,
		JoinedAt:    timestamppb.New(i.JoinedAt),
		RequestedAt: timestamppb.New(i.RequestedAt),
	}
}
