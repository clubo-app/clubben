package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/participation-service/repository"
	"github.com/clubo-app/clubben/protobuf/participation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *server) GetPartyParticipant(ctx context.Context, req *participation.UserPartyRequest) (*participation.PartyParticipant, error) {
	p, err := s.pp.GetPartyParticipant(ctx, repository.UserPartyParams{
		UserId:  req.UserId,
		PartyId: req.PartyId,
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}
	if p.PartyId == "" && p.UserId == "" {
		return nil, status.Error(codes.NotFound, "Participant not found")
	}

	return p.ToGRPCPartyParticipant(), nil
}
