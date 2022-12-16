package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/relation"
)

func (r *relationServer) GetFavoritePartyManyParties(ctx context.Context, req *relation.GetFavoritePartyManyPartiesRequest) (*relation.ManyFavoritePartiesMap, error) {
	favoriteParties, err := r.fp.GetFavoritePartyManyParties(ctx, req.UserId, req.PartyIds)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	res := make(map[string]*relation.FavoriteParty)
	for _, fp := range favoriteParties {
		res[fp.PartyId] = fp.ToGRPCFavoriteParty()
	}

	return &relation.ManyFavoritePartiesMap{FavoriteParties: res}, nil
}
