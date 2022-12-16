package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/relation"
)

func (r *relationServer) GetFavoritePartyManyUser(ctx context.Context, req *relation.GetFavoritePartyManyUserRequest) (*relation.ManyFavoritePartiesMap, error) {
	favoriteParties, err := r.fp.GetFavoritePartyManyUser(ctx, req.UserIds, req.PartyId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	res := make(map[string]*relation.FavoriteParty)
	for _, fp := range favoriteParties {
		res[fp.UserId] = fp.ToGRPCFavoriteParty()
	}

	return &relation.ManyFavoritePartiesMap{FavoriteParties: res}, nil
}
