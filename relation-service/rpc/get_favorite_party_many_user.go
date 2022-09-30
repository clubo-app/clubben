package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *relationServer) GetFavoritePartyManyUser(ctx context.Context, req *relation.GetFavoritePartyManyUserRequest) (*relation.ManyFavoritePartiesMap, error) {
	_, err := ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	favoriteParties, err := s.fp.GetFavoritePartyManyUser(ctx, req.UserIds, req.PartyId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	res := make(map[string]*relation.FavoriteParty)
	for _, fp := range favoriteParties {
		res[fp.UserId] = fp.ToGRPCFavoriteParty()
	}

	return &relation.ManyFavoritePartiesMap{FavoriteParties: res}, nil
}
