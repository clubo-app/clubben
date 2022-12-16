package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/participation-service/service"
	"github.com/clubo-app/clubben/protobuf/participation"
)

func (s *server) JoinParty(ctx context.Context, req *participation.UserPartyRequest) (*participation.PartyParticipant, error) {
	p, err := s.pp.Join(ctx, service.JoinParams{
		UserId:  req.UserId,
		PartyId: req.PartyId,
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCPartyParticipant(), nil
}
