package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/participation-service/repository"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) AcceptPartyInvite(ctx context.Context, req *participation.DeclinePartyInviteRequest) (*participation.PartyParticipant, error) {
	_, err := ksuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}
	_, err = ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	p, err := s.pi.Accept(ctx, repository.DeclineParams{
		UserId:    req.UserId,
		PartyId:   req.PartyId,
		InviterId: req.InviterId,
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCPartyParticipant(), nil
}
