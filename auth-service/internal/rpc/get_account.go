package rpc

import (
	"context"

	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	"github.com/clubo-app/clubben/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) GetAccount(ctx context.Context, req *pbauth.GetAccountRequest) (*pbauth.Account, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}

	a, err := s.accountService.GetById(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return a.ToGRPCAccount(), nil
}
