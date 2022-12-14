package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	cg "github.com/clubo-app/clubben/protobuf/common"
	"github.com/clubo-app/clubben/protobuf/events"
	rg "github.com/clubo-app/clubben/protobuf/relation"
)

func (s *relationServer) RemoveFriend(ctx context.Context, req *rg.RemoveFriendRequest) (*cg.SuccessIndicator, error) {
	err := s.fs.RemoveFriendRelation(ctx, req.UserId, req.FriendId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	s.stream.PublishEvent(&events.FriendRemoved{
		UserId:   req.UserId,
		FriendId: req.FriendId,
	})

	return &cg.SuccessIndicator{Sucess: true}, nil
}
