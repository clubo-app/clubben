package datastruct

import (
	"time"

	"github.com/clubo-app/protobuf/participation"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Invite struct {
	UserId     string    `db:"user_id"     validate:"required"`
	InviterId  string    `db:"inviter_id"  validate:"required"`
	PartyId    string    `db:"party_id"    validate:"required"`
	ValidUntil time.Time `db:"valid_until" validate:"required"`
}

func (i Invite) ToGRPCPartyInvite() *participation.PartyInvite {
	return &participation.PartyInvite{
		UserId:     i.UserId,
		PartyId:    i.PartyId,
		InviterId:  i.InviterId,
		ValidUntil: timestamppb.New(i.ValidUntil),
	}
}
