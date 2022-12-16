package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/participation-service/repository"
	"github.com/clubo-app/clubben/protobuf/participation"
)

func (s *server) AcceptPartyInvite(ctx context.Context, req *participation.DeclinePartyInviteRequest) (*participation.PartyParticipant, error) {
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
