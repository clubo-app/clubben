package datastruct

import (
	"time"

	"github.com/clubo-app/protobuf/participation"
)

type AggregatedPartyInvite struct {
	User       *AggregatedProfile `json:"profile,omitempty"`
	Inviter    AggregatedProfile  `json:"inviter,omitempty"`
	Party      *AggregatedParty   `json:"party,omitempty"`
	ValidUntil string             `json:"valid_until"`
}
type PagedAggregatedPartyInvite struct {
	Invites  []AggregatedPartyInvite `json:"invites"`
	NextPage string                  `json:"nextPage"`
}

func PartyInviteToAgg(i *participation.PartyInvite) AggregatedPartyInvite {
	if i == nil {
		return AggregatedPartyInvite{}
	}
	return AggregatedPartyInvite{
		ValidUntil: i.ValidUntil.AsTime().UTC().Format(time.RFC3339),
	}
}

func (i AggregatedPartyInvite) AddUser(u AggregatedProfile) AggregatedPartyInvite {
	if u.Id != "" {
		i.User = &u
	}
	return i
}
func (i AggregatedPartyInvite) AddInviter(u AggregatedProfile) AggregatedPartyInvite {
	if u.Id != "" {
		i.Inviter = u
	}
	return i
}
func (i AggregatedPartyInvite) AddParty(p AggregatedParty) AggregatedPartyInvite {
	if p.Id != "" {
		i.Party = &p
	}
	return i
}

type AggregatedPartyParticipant struct {
	User        AggregatedProfile `json:"user"`
	Party       AggregatedParty   `json:"party"`
	Requested   bool              `json:"requested"`
	JoinedAt    string            `json:"joined_at"`
	RequestedAt string            `json:"requested_at"`
}

type PagedAggregatedPartyParticipant struct {
	Participants []AggregatedPartyParticipant `json:"participants"`
	NextPage     string                       `json:"nextPage"`
}

func PartyParticipantToAgg(p *participation.PartyParticipant) AggregatedPartyParticipant {
	if p == nil {
		return AggregatedPartyParticipant{}
	}
	return AggregatedPartyParticipant{
		JoinedAt:    p.JoinedAt.AsTime().UTC().Format(time.RFC3339),
		RequestedAt: p.RequestedAt.AsTime().UTC().Format(time.RFC3339),
		Requested:   p.Requested,
	}
}
func (p AggregatedPartyParticipant) AddUser(u AggregatedProfile) AggregatedPartyParticipant {
	if u.Id != "" {
		p.User = u
	}
	return p
}
func (pp AggregatedPartyParticipant) AddParty(p AggregatedParty) AggregatedPartyParticipant {
	if p.Id != "" {
		pp.Party = p
	}
	return pp
}
