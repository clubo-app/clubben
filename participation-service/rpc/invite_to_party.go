package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/participation-service/repository"
	"github.com/clubo-app/clubben/protobuf/participation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) InviteToParty(ctx context.Context, req *participation.InviteToPartyRequest) (*participation.PartyInvite, error) {
	if req.UserId == req.InviterId {
		return nil, status.Error(codes.InvalidArgument, "You can't invite yourself")
	}

	i, err := s.pi.Invite(ctx, repository.InviteParams{
		UserId:    req.UserId,
		InviterId: req.InviterId,
		PartyId:   req.PartyId,
		ValidFor:  req.ValidFor.AsDuration(),
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return i.ToGRPCPartyInvite(), nil
}
