package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/search"
	"github.com/clubo-app/clubben/search-service/datastruct"
)

func (s *searchServer) SearchParties(ctx context.Context, req *search.SearchPartiesRequest) (*search.SearchPartiesResponse, error) {
	parties, err := s.party.QueryParty(ctx, req.Query, datastruct.Location{Lat: req.Lat, Lng: req.Long})
	if err != nil {
		return &search.SearchPartiesResponse{}, utils.HandleError(err)
	}

	res := make([]*search.IndexedParty, len(parties))
	for i, p := range parties {
		res[i] = p.ToGRPCParty()
	}

	return &search.SearchPartiesResponse{
		Parties: res,
	}, nil
}
