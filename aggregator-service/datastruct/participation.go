package datastruct

import (
	"time"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/participation"
)

type ParticipationStatus struct {
	Requested     bool   `json:"requested"`
	Participating bool   `json:"participating"`
	RequestedAt   string `json:"requested_at,omitempty"`
	JoinedAt      string `json:"joined_at,omitempty"`
}

// This parses a *participation.PartyParticipant to a ParticipationStatus.
// The proto Message can be nil and would just return default ParticipationStatus.
// The ParticipationStatus holds weither the User is participating/requsting the Party.
func ParseParticipationStatus(p *participation.PartyParticipant) *ParticipationStatus {
	if p == nil {
		return nil
	}

	s := new(ParticipationStatus)
	s.Requested = p.Requested
	if !p.Requested {
		s.Participating = true
	}

	if !utils.TimestamppIsZero(p.JoinedAt) {
		s.JoinedAt = p.JoinedAt.AsTime().UTC().Format(time.RFC3339)
	}
	if !utils.TimestamppIsZero(p.RequestedAt) {
		s.RequestedAt = p.RequestedAt.AsTime().UTC().Format(time.RFC3339)
	}
	return s
}

type AggregatedPartyParticipant struct {
	Profile     *AggregatedProfile `json:"profile"`
	Party       *AggregatedParty   `json:"party"`
	Requested   bool               `json:"requested"`
	JoinedAt    string             `json:"joined_at"`
	RequestedAt string             `json:"requested_at"`
}

type PagedAggregatedPartyParticipant struct {
	Participants []AggregatedPartyParticipant `json:"participants"`
	NextPage     string                       `json:"next_page"`
}

func PartyParticipantToAgg(p *participation.PartyParticipant) *AggregatedPartyParticipant {
	if p == nil {
		return &AggregatedPartyParticipant{}
	}
	agg := AggregatedPartyParticipant{
		Requested: p.Requested,
	}
	if !utils.TimestamppIsZero(p.JoinedAt) {
		agg.JoinedAt = p.JoinedAt.AsTime().UTC().Format(time.RFC3339)
	}
	if !utils.TimestamppIsZero(p.RequestedAt) {
		agg.RequestedAt = p.RequestedAt.AsTime().UTC().Format(time.RFC3339)
	}
	return &agg
}

func (p *AggregatedPartyParticipant) AddProfile(u *AggregatedProfile) *AggregatedPartyParticipant {
	if u.Id != "" {
		p.Profile = u
	}
	return p
}

func (pp *AggregatedPartyParticipant) AddParty(p *AggregatedParty) *AggregatedPartyParticipant {
	if p.Id != "" {
		pp.Party = p
	}
	return pp
}
