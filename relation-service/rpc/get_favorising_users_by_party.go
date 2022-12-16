package rpc

import (
	"context"
	"encoding/base64"

	"github.com/clubo-app/clubben/libs/utils"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *relationServer) GetFavorisingUsersByParty(ctx context.Context, req *rg.GetFavorisingUsersByPartyRequest) (*rg.PagedFavoriteParties, error) {
	p, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}

	fps, p, err := s.fp.GetFavorisingUsersByParty(ctx, req.PartyId, p, req.Limit)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	nextPage := base64.URLEncoding.EncodeToString(p)

	res := make([]*rg.FavoriteParty, len(fps))
	for i, fp := range fps {
		res[i] = fp.ToGRPCFavoriteParty()
	}

	return &rg.PagedFavoriteParties{FavoriteParties: res, NextPage: nextPage}, nil
}
