package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	rg "github.com/clubo-app/clubben/protobuf/relation"
)

func (s *relationServer) CreateFriendRequest(ctx context.Context, req *rg.CreateFriendRequestRequest) (*rg.FriendRelation, error) {
	fr, err := s.fs.CreateFriendRequest(ctx, req.UserId, req.FriendId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return fr.ToGRPCFriendRelation(), nil
}
