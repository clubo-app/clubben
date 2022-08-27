package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/participation-service/service"
	"github.com/clubo-app/protobuf/participation"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s server) JoinParty(ctx context.Context, req *participation.UserPartyRequest) (*participation.PartyParticipant, error) {
	_, err := ksuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}
	_, err = ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	p, err := s.pp.Join(ctx, service.JoinParams{
		UserId:  req.UserId,
		PartyId: req.PartyId,
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCPartyParticipant(), nil
}
