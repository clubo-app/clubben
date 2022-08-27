package rpc

import (
	"context"
	"encoding/base64"

	"github.com/clubo-app/packages/utils"
	rg "github.com/clubo-app/protobuf/relation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s relationServer) GetFriends(ctx context.Context, req *rg.GetFriendsRequest) (*rg.PagedFriendRelations, error) {
	p, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}

	fs, p, err := s.fs.GetFriends(ctx, req.UserId, p, req.Limit)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	nextPage := base64.URLEncoding.EncodeToString(p)

	res := make([]*rg.FriendRelation, len(fs))
	for i, fr := range fs {
		res[i] = fr.ToGRPCFriendRelation()
	}

	return &rg.PagedFriendRelations{Relations: res, NextPage: nextPage}, nil
}
