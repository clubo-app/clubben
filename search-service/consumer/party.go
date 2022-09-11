package consumer

import (
	"context"
	"log"

	"github.com/clubo-app/clubben/protobuf/events"
	"github.com/clubo-app/clubben/search-service/datastruct"
	"github.com/clubo-app/clubben/search-service/repository"
)

type partyConsumer struct {
	repo *repository.PartyRepository
}

func NewPartyConsumer(repo *repository.PartyRepository) partyConsumer {
	return partyConsumer{
		repo: repo,
	}
}

func (c partyConsumer) PartyCreated(p *events.PartyCreated) {
	if p == nil {
		return
	}

	err := c.repo.PutParty(context.Background(), datastruct.Party{
		Id:            p.Party.Id,
		Title:         p.Party.Title,
		Description:   p.Party.Description,
		MusicGenre:    p.Party.MusicGenre,
		Location:      datastruct.Location{Lat: p.Party.Lat, Lng: p.Party.Long},
		EntryDate:     p.Party.EntryDate.AsTime().Unix(),
		FavoriteCount: 0,
	})
	if err != nil {
		log.Println("Error inserting Party: ", err)
	}
	log.Printf("Party Created %+v", p)
}

func (c partyConsumer) PartyUpdated(p *events.PartyUpdated) {
	if p == nil {
		return
	}

	err := c.repo.UpdateParty(context.Background(), datastruct.Party{
		Id:          p.Party.Id,
		Title:       p.Party.Title,
		Description: p.Party.Description,
		MusicGenre:  p.Party.MusicGenre,
		Location:    datastruct.Location{Lat: p.Party.Lat, Lng: p.Party.Long},
		EntryDate:   p.Party.EntryDate.AsTime().Unix(),
	})
	if err != nil {
		log.Println("Error inserting Party: ", err)
	}
	log.Printf("Party Created %+v", p)
}

func (c partyConsumer) PartyFavorited(e *events.PartyFavorited) {
	if e == nil {
		return
	}
	err := c.repo.IncrementFavoriteCount(context.Background(), e.PartyId)
	if err != nil {
		log.Println("Error incrementing FavoriteCount: ", err)
	}
	log.Println("FavoriteCount incremented for Party ", e.PartyId)
}

func (c partyConsumer) PartyUnfavorited(e *events.PartyUnfavorited) {
	if e == nil {
		return
	}
	err := c.repo.DecrementFavoriteCount(context.Background(), e.PartyId)
	if err != nil {
		log.Println("Error decrementing FavoriteCount: ", err)
	}
	log.Println("FavoriteCount decremented for Party ", e.PartyId)
}
