package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	cg "github.com/clubo-app/protobuf/common"
	rg "github.com/clubo-app/protobuf/relation"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s relationServer) CreateFriendRequest(ctx context.Context, req *rg.CreateFriendRequestRequest) (*cg.SuccessIndicator, error) {
	_, err := ksuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}
	_, err = ksuid.Parse(req.FriendId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Friend id")
	}

	err = s.fs.CreateFriendRequest(ctx, req.UserId, req.FriendId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &cg.SuccessIndicator{Sucess: true}, nil
}
