package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *relationServer) GetFavoritePartyManyParties(ctx context.Context, req *relation.GetFavoritePartyManyPartiesRequest) (*relation.ManyFavoritePartiesMap, error) {
	_, err := ksuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}

	favoriteParties, err := r.fp.GetFavoritePartyManyParties(ctx, req.UserId, req.PartyIds)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	res := make(map[string]*relation.FavoriteParty)
	for _, fp := range favoriteParties {
		res[fp.UserId] = fp.ToGRPCFavoriteParty()
	}

	return &relation.ManyFavoritePartiesMap{FavoriteParties: res}, nil
}
