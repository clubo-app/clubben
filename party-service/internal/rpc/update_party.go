package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/party-service/internal/dto"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	"github.com/paulmach/orb"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s partyServer) UpdateParty(c context.Context, req *pbparty.UpdatePartyRequest) (*pbparty.Party, error) {
	id, err := ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	p := dto.Party{
		ID:            id.String(),
		UserId:        req.RequesterId,
		Title:         req.Title,
		Description:   req.Description,
		MusicGenre:    req.MusicGenre,
		Location:      orb.Point{float64(req.Long), float64(req.Lat)},
		StreetAddress: req.StreetAddress,
		PostalCode:    req.PostalCode,
		State:         req.State,
		Country:       req.Country,
		EntryDate:     req.EntryDate.AsTime(),
	}

	d, err := s.ps.Update(c, p)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return d.ToGRPCParty(), nil
}
