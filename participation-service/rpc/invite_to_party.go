package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/participation-service/repository"
	"github.com/clubo-app/protobuf/participation"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s server) InviteToParty(ctx context.Context, req *participation.InviteToPartyRequest) (*participation.PartyInvite, error) {
	_, err := ksuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}
	_, err = ksuid.Parse(req.InviterId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Inviter id")
	}
	_, err = ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

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
