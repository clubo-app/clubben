package datastruct

import (
	"time"

	pg "github.com/clubo-app/protobuf/party"
	"github.com/clubo-app/protobuf/profile"
	sg "github.com/clubo-app/protobuf/story"
)

type AggregatedParty struct {
	Id              string             `json:"id,omitempty"`
	Creator         *AggregatedProfile `json:"creator,omitempty"`
	Title           string             `json:"title,omitempty"`
	IsPublic        bool               `json:"is_public"`
	MaxParticipants uint32             `json:"max_participants"`
	Lat             float32            `json:"lat,omitempty"`
	Long            float32            `json:"long,omitempty"`
	StreetAddress   string             `json:"street_address,omitempty"`
	PostalCode      string             `json:"postal_code,omitempty"`
	State           string             `json:"state,omitempty"`
	Country         string             `json:"country,omitempty"`
	Stories         []*sg.Story        `json:"stories,omitempty"`
	StartDate       string             `json:"start_date,omitempty"`
	CreatedAt       string             `json:"created_at,omitempty"`
	FavoriteCount   uint32             `json:"favorite_count"`
}

func PartyToAgg(p *pg.Party) AggregatedParty {
	if p == nil {
		return AggregatedParty{}
	}
	return AggregatedParty{
		Id:              p.Id,
		Title:           p.Title,
		IsPublic:        p.IsPublic,
		MaxParticipants: p.MaxParticipants,
		Lat:             p.Lat,
		Long:            p.Long,
		StreetAddress:   p.StreetAddress,
		Stories:         []*sg.Story{},
		PostalCode:      p.PostalCode,
		State:           p.State,
		Country:         p.Country,
		StartDate:       p.StartDate.AsTime().UTC().Format(time.RFC3339),
		CreatedAt:       p.CreatedAt.AsTime().UTC().Format(time.RFC3339),
	}
}

func (p AggregatedParty) AddCreator(prof AggregatedProfile) AggregatedParty {
	if prof.Id != "" {
		p.Creator = &prof
	}
	return p
}

func (p AggregatedParty) AddStory(s []*sg.Story) AggregatedParty {
	if s != nil {
		p.Stories = s
	}
	return p
}

func (p AggregatedParty) AddFCount(c uint32) AggregatedParty {
	p.FavoriteCount = c
	return p
}

type PagedAggregatedParty struct {
	Parties []AggregatedParty `json:"parties"`
}

type AggregatedFavoriteParty struct {
	UserId      string          `json:"user_id"`
	Party       AggregatedParty `json:"party"`
	FavoritedAt string          `json:"favorited_at"`
}

type AggregatedFavorisingUsers struct {
	User        *profile.Profile `json:"user"`
	PartyId     string           `json:"party_id"`
	FavoritedAt string           `json:"favorited_at"`
}

type PagedAggregatedFavoriteParty struct {
	FavoriteParties []AggregatedFavoriteParty `json:"favorite_parties"`
	NextPage        string                    `json:"nextPage"`
}

type PagedAggregatedFavorisingUsers struct {
	FavoriteParties []AggregatedFavorisingUsers `json:"favorite_parties"`
	NextPage        string                      `json:"nextPage"`
}
