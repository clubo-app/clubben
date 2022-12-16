package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/relation"
)

func (s *relationServer) GetFavoriteParty(ctx context.Context, req *relation.PartyAndUserRequest) (*relation.FavoriteParty, error) {
	fp, err := s.fp.GetFavoriteParty(ctx, req.UserId, req.PartyId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return fp.ToGRPCFavoriteParty(), nil
}
