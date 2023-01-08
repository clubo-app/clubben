package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/party-service/internal/repository"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
)

func (s partyServer) GeoSearch(ctx context.Context, req *pbparty.GeoSearchRequest) (*pbparty.PagedParties, error) {
	ps, err := s.ps.GeoSearch(ctx, repository.GeoSearchParams{
		Lat:            req.Lat,
		Long:           req.Long,
		RadiusInDegree: req.RadiusInDegrees,
		IsPublic:       req.IsPublic,
		Offset:         req.Offset,
		Limit:          req.Limit,
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	var pp []*pbparty.Party
	for _, p := range ps {
		pp = append(pp, p.ToGRPCParty())
	}

	return &pbparty.PagedParties{Parties: pp}, nil
}
