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

func (s *server) GetUserInvites(ctx context.Context, req *participation.GetUserInvitesRequest) (*participation.PagedPartyInvites, error) {
	p, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}

	_, err = ksuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}

	pp, p, err := s.pi.GetUserInvites(ctx, repository.GetUserInvitesParams{
		UId:   req.UserId,
		Page:  p,
		Limit: int(req.Limit),
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	nextPage := base64.URLEncoding.EncodeToString(p)

	res := make([]*participation.PartyInvite, len(pp))
	for i, pr := range pp {
		res[i] = pr.ToGRPCPartyInvite()
	}

	return &participation.PagedPartyInvites{Invites: res, NextPage: nextPage}, nil
}
