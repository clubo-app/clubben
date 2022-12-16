package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	cg "github.com/clubo-app/clubben/protobuf/common"
	rg "github.com/clubo-app/clubben/protobuf/relation"
)

func (s *relationServer) AcceptFriendRequest(ctx context.Context, req *rg.AcceptFriendRequestRequest) (*cg.SuccessIndicator, error) {
	err := s.fs.AcceptFriendRequest(ctx, req.UserId, req.FriendId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &cg.SuccessIndicator{Sucess: true}, nil
}
