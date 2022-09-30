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

func (s *server) GetPartyRequests(ctx context.Context, req *participation.GetPartyParticipantsRequest) (*participation.PagedPartyParticipants, error) {
	p, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}

	_, err = ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	pp, p, err := s.pp.GetPartyRequests(ctx, repository.GetPartyParticipantsParams{
		PId:   req.PartyId,
		Page:  p,
		Limit: int(req.Limit),
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	nextPage := base64.URLEncoding.EncodeToString(p)

	res := make([]*participation.PartyParticipant, len(pp))
	for i, pr := range pp {
		res[i] = pr.ToGRPCPartyParticipant()
	}

	return &participation.PagedPartyParticipants{Participants: res, NextPage: nextPage}, nil

}
