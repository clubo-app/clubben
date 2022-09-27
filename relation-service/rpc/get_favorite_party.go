package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s relationServer) GetFavoriteParty(ctx context.Context, req *relation.PartyAndUserRequest) (*relation.FavoriteParty, error) {
	_, err := ksuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}
	_, err = ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	fp, err := s.fp.GetFavoriteParty(ctx, req.UserId, req.PartyId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return fp.ToGRPCFavoriteParty(), nil
}
