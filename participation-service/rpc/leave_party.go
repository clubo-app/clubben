package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/participation-service/repository"
	"github.com/clubo-app/clubben/protobuf/participation"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) LeaveParty(ctx context.Context, req *participation.UserPartyRequest) (*emptypb.Empty, error) {
	err := s.pp.Leave(ctx, repository.UserPartyParams{
		UserId:  req.UserId,
		PartyId: req.PartyId,
	})
	if err != nil {
		return &emptypb.Empty{}, utils.HandleError(err)
	}

	return &emptypb.Empty{}, nil
}
