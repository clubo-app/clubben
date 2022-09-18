package rpc

import (
	"context"
	"encoding/base64"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/participation-service/repository"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s server) GetUserParticipation(ctx context.Context, req *participation.GetUserParticipationRequest) (*participation.PagedPartyParticipants, error) {
	page, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}

	_, err = ksuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	participations, nextPage, err := s.pp.GetUserParticipation(ctx, repository.GetUserParticipationParams{
		UId:   req.UserId,
		Page:  page,
		Limit: int(req.Limit),
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	res := make([]*participation.PartyParticipant, len(participations))
	for i, pr := range participations {
		res[i] = pr.ToGRPCPartyParticipant()
	}

	return &participation.PagedPartyParticipants{
		Participants: res,
		NextPage:     base64.URLEncoding.EncodeToString(nextPage),
	}, nil
}
