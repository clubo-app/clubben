package rpc

import (
	"context"

	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	"github.com/clubo-app/clubben/libs/utils"
)

func (s *authServer) GetAccount(ctx context.Context, req *pbauth.GetAccountRequest) (*pbauth.Account, error) {
	a, err := s.accountService.GetById(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return a.ToGRPCAccount(), nil
}
