package datastruct

import (
	"time"

	"github.com/clubo-app/clubben/protobuf/participation"
)

type AggregatedPartyInvite struct {
	Profile    *AggregatedProfile `json:"profile,omitempty"`
	Inviter    AggregatedProfile  `json:"inviter,omitempty"`
	Party      *AggregatedParty   `json:"party,omitempty"`
	ValidUntil string             `json:"valid_until"`
}
type PagedAggregatedPartyInvite struct {
	Invites  []AggregatedPartyInvite `json:"invites"`
	NextPage string                  `json:"next_page"`
}

func PartyInviteToAgg(i *participation.PartyInvite) AggregatedPartyInvite {
	if i == nil {
		return AggregatedPartyInvite{}
	}
	return AggregatedPartyInvite{
		ValidUntil: i.ValidUntil.AsTime().UTC().Format(time.RFC3339),
	}
}

func (i AggregatedPartyInvite) AddProfile(u AggregatedProfile) AggregatedPartyInvite {
	if u.Id != "" {
		i.Profile = &u
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
