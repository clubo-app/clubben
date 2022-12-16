package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	cg "github.com/clubo-app/clubben/protobuf/common"
	rg "github.com/clubo-app/clubben/protobuf/relation"
)

func (s *relationServer) DeclineFriendRequest(ctx context.Context, req *rg.DeclineFriendRequestRequest) (*cg.SuccessIndicator, error) {
	err := s.fs.DeclineFriendRequest(ctx, req.UserId, req.FriendId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &cg.SuccessIndicator{Sucess: true}, nil
}
