package rpc

import (
	"context"

	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	"github.com/clubo-app/clubben/libs/id"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *authServer) RegisterAnonymously(ctx context.Context, req *emptypb.Empty) (*pbauth.Account, error) {
	a, err := s.accountService.CreateAnonymously(ctx, id.New(id.User))
	if err != nil {
		return nil, err
	}

	return a.ToGRPCAccount(), nil
}
