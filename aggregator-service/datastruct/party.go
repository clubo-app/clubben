package datastruct

import (
	"log"
	"time"

	"github.com/clubo-app/clubben/libs/utils"
	pg "github.com/clubo-app/clubben/protobuf/party"
	sg "github.com/clubo-app/clubben/protobuf/story"
)

type AggregatedParty struct {
	Id                  string               `json:"id"`
	Creator             *AggregatedProfile   `json:"creator,omitempty"`
	Title               string               `json:"title,omitempty"`
	Description         string               `json:"description,omitempty"`
	IsPublic            bool                 `json:"is_public"`
	MusicGenre          string               `json:"music_genre,omitempty"`
	MaxParticipants     int32                `json:"max_participants"`
	ParticipationStatus *ParticipationStatus `json:"particpation_status,omitempty"`
	Lat                 float32              `json:"lat,omitempty"`
	Lon                 float32              `json:"lon,omitempty"`
	StreetAddress       string               `json:"street_address,omitempty"`
	PostalCode          string               `json:"postal_code,omitempty"`
	State               string               `json:"state,omitempty"`
	Country             string               `json:"country,omitempty"`
	Stories             []*sg.Story          `json:"stories,omitempty"`
	EntryDate           string               `json:"entry_date,omitempty"`
	CreatedAt           string               `json:"created_at,omitempty"`
	FavoriteCount       uint32               `json:"favorite_count"`
}

func PartyToAgg(p *pg.Party) *AggregatedParty {
	if p == nil {
		return &AggregatedParty{}
	}
	agg := &AggregatedParty{
		Id:              p.Id,
		Title:           p.Title,
		Description:     p.Description,
		IsPublic:        p.IsPublic,
		MusicGenre:      p.MusicGenre,
		MaxParticipants: p.MaxParticipants,
		Lat:             p.Lat,
		Lon:             p.Long,
		StreetAddress:   p.StreetAddress,
		Stories:         []*sg.Story{},
		PostalCode:      p.PostalCode,
		State:           p.State,
		Country:         p.Country,
	}
	if !utils.TimestamppIsZero(p.EntryDate) {
		log.Println("Adding EntryDate")
		agg.EntryDate = p.EntryDate.AsTime().UTC().Format(time.RFC3339)
	}
	if !utils.TimestamppIsZero(p.CreatedAt) {
		log.Println("Adding CreatedAt Date")
		agg.CreatedAt = p.CreatedAt.AsTime().UTC().Format(time.RFC3339)
	}

	return agg
}

func (p *AggregatedParty) AddCreator(prof *AggregatedProfile) *AggregatedParty {
	if prof.Id != "" {
		p.Creator = prof
	}
	return p
}

func (p *AggregatedParty) AddStory(s []*sg.Story) *AggregatedParty {
	if s != nil {
		p.Stories = s
	}
	return p
}

func (p *AggregatedParty) AddFCount(c uint32) *AggregatedParty {
	p.FavoriteCount = c
	return p
}

func (p *AggregatedParty) AddParticipationStatus(s ParticipationStatus) *AggregatedParty {
	p.ParticipationStatus = &s
	return p
}

type PagedAggregatedParty struct {
	Parties []*AggregatedParty `json:"parties"`
}

type AggregatedFavoriteParty struct {
	UserId      string           `json:"user_id"`
	Party       *AggregatedParty `json:"party"`
	FavoritedAt string           `json:"favorited_at"`
}

type PagedAggregatedFavoriteParty struct {
	FavoriteParties []AggregatedFavoriteParty `json:"favorite_parties"`
	NextPage        string                    `json:"next_page"`
}
