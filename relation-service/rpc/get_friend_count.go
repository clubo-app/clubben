package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	rg "github.com/clubo-app/clubben/protobuf/relation"
)

func (s *relationServer) GetFriendCount(ctx context.Context, req *rg.GetFriendCountRequest) (*rg.GetFriendCountResponse, error) {
	fc, err := s.fs.GetFriendCount(ctx, req.UserId)
	if err != nil {
		return &rg.GetFriendCountResponse{FriendCount: 0}, utils.HandleError(err)
	}

	return &rg.GetFriendCountResponse{FriendCount: uint32(fc.FriendCount)}, nil
}
