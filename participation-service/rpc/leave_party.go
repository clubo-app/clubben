package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/participation-service/repository"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) LeaveParty(ctx context.Context, req *participation.UserPartyRequest) (*emptypb.Empty, error) {
	_, err := ksuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}
	_, err = ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	err = s.pp.Leave(ctx, repository.UserPartyParams{
		UserId:  req.UserId,
		PartyId: req.PartyId,
	})
	if err != nil {
		return &emptypb.Empty{}, utils.HandleError(err)
	}

	return &emptypb.Empty{}, nil
}
