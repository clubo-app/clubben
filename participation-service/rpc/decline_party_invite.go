package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/participation-service/repository"
	"github.com/clubo-app/clubben/protobuf/participation"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) DeclinePartyInvite(ctx context.Context, req *participation.DeclinePartyInviteRequest) (*emptypb.Empty, error) {
	err := s.pi.Decline(ctx, repository.DeclineParams{
		UserId:    req.UserId,
		InviterId: req.InviterId,
		PartyId:   req.PartyId,
	})
	if err != nil {
		return &emptypb.Empty{}, utils.HandleError(err)
	}

	return &emptypb.Empty{}, nil
}
