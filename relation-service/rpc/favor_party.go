package rpc

import (
	"context"
	"time"

	"github.com/clubo-app/clubben/libs/utils"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/clubo-app/clubben/relation-service/datastruct"
)

func (s *relationServer) FavorParty(ctx context.Context, req *rg.PartyAndUserRequest) (*rg.FavoriteParty, error) {
	fp, err := s.fp.FavorParty(ctx, datastruct.FavoriteParty{
		UserId:      req.UserId,
		PartyId:     req.PartyId,
		FavoritedAt: time.Now(),
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return fp.ToGRPCFavoriteParty(), nil
}
