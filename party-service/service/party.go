package service

import (
	"context"
	"database/sql"

	"github.com/clubo-app/clubben/libs/stream"
	"github.com/clubo-app/clubben/party-service/dto"
	"github.com/clubo-app/clubben/party-service/repository"
	"github.com/clubo-app/clubben/protobuf/events"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PartyService interface {
	Create(ctx context.Context, p dto.Party) (repository.Party, error)
	Update(ctx context.Context, p dto.Party) (repository.Party, error)
	Delete(ctx context.Context, uId, pId string) error
	Get(ctx context.Context, pId string) (repository.Party, error)
	GetMany(ctx context.Context, ids []string) ([]repository.Party, error)
	GetByUser(ctx context.Context, params repository.GetPartiesByUserParams) ([]repository.Party, error)
	GeoSearch(ctx context.Context, params repository.GeoSearchParams) ([]repository.Party, error)
}

type partyService struct {
	r      *repository.PartyRepository
	stream stream.Stream
}

func NewPartyService(r *repository.PartyRepository, stream stream.Stream) PartyService {
	return &partyService{
		r:      r,
		stream: stream,
	}
}

func (s partyService) Create(ctx context.Context, p dto.Party) (res repository.Party, err error) {
	res, err = s.r.CreateParty(ctx, repository.CreatePartyParams{
		ID:              ksuid.New().String(),
		UserID:          p.UserId,
		Title:           p.Title,
		Description:     sql.NullString{Valid: p.Description != "", String: p.Description},
		IsPublic:        p.IsPublic,
		MusicGenre:      p.MusicGenre,
		MaxParticipants: p.MaxParticipants,
		Location:        p.Location,
		StreetAddress:   p.StreetAddress,
		PostalCode:      p.PostalCode,
		State:           p.State,
		Country:         p.Country,
		EntryDate:       p.EntryDate,
	})
	if err != nil {
		return res, err
	}

	s.stream.PublishEvent(&events.PartyCreated{Party: res.ToGRPCParty()})

	return res, nil
}

func (s partyService) Update(ctx context.Context, p dto.Party) (res repository.Party, err error) {
	res, err = s.r.UpdateParty(ctx, repository.UpdatePartyParams{
		ID:            p.ID,
		Title:         p.Title,
		Description:   p.Description,
		MusicGenre:    p.MusicGenre,
		Location:      p.Location,
		StreetAddress: p.StreetAddress,
		PostalCode:    p.PostalCode,
		State:         p.State,
		Country:       p.Country,
		EntryDate:     p.EntryDate,
	})
	if err != nil {
		return res, err
	}

	updatedValues := party.Party{
		Id:              p.ID,
		UserId:          p.UserId,
		Title:           p.Title,
		Description:     p.Description,
		MusicGenre:      p.MusicGenre,
		MaxParticipants: p.MaxParticipants,
		StreetAddress:   p.StreetAddress,
		PostalCode:      p.PostalCode,
		State:           p.State,
		Country:         p.Country,
	}

	if p.Location.Lat() != 0 && p.Location.Lon() != 0 {
		updatedValues.Lat = float32(p.Location.Lat())
		updatedValues.Long = float32(p.Location.Lon())
	}
	entryYear := p.EntryDate.Year()
	if !(entryYear == 1970) {
		updatedValues.EntryDate = timestamppb.New(p.EntryDate)
	}

	s.stream.PublishEvent(&events.PartyUpdated{Party: &updatedValues})

	return res, nil
}

func (s partyService) Delete(ctx context.Context, uId, pId string) error {
	err := s.r.DeleteParty(ctx, repository.DeletePartyParams{
		ID:     pId,
		UserID: uId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s partyService) Get(ctx context.Context, pId string) (res repository.Party, err error) {
	res, err = s.r.GetParty(ctx, pId)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s partyService) GetMany(ctx context.Context, ids []string) (res []repository.Party, err error) {
	res, err = s.r.GetManyParties(ctx, repository.GetManyPartiesParams{
		Ids:   ids,
		Limit: int32(len(ids)),
	})
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s partyService) GetByUser(ctx context.Context, params repository.GetPartiesByUserParams) (res []repository.Party, err error) {
	res, err = s.r.GetPartiesByUser(ctx, params)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s partyService) GeoSearch(ctx context.Context, params repository.GeoSearchParams) ([]repository.Party, error) {
	res, err := s.r.GeoSearch(ctx, params)
	if err != nil {
		return res, err
	}

	return res, nil
}
