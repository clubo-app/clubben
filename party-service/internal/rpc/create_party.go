package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/party-service/internal/dto"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	"github.com/paulmach/orb"
)

func (s partyServer) CreateParty(c context.Context, req *pbparty.CreatePartyRequest) (*pbparty.Party, error) {
	d := dto.Party{
		Title:           req.Title,
		Description:     req.Description,
		UserId:          req.RequesterId,
		Location:        orb.Point{float64(req.Long), float64(req.Lat)},
		IsPublic:        req.IsPublic,
		MusicGenre:      req.MusicGenre,
		MaxParticipants: req.MaxParticipants,
		StreetAddress:   req.StreetAddress,
		PostalCode:      req.PostalCode,
		State:           req.State,
		Country:         req.Country,
		EntryDate:       req.EntryDate.AsTime(),
	}

	p, err := s.ps.Create(c, d)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCParty(), nil
}
