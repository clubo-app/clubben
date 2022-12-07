package rpc

import (
	"context"

	"github.com/clubo-app/clubben/auth-service/internal/repository"
	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	"github.com/clubo-app/clubben/libs/utils"
)

func (s *authServer) UpdateAccount(ctx context.Context, req *pbauth.UpdateAccountRequest) (*pbauth.Account, error) {
	params := repository.UpdateAccountParams{
		UId:      req.Id,
		Password: req.Password,
		Email:    req.Email,
	}

	if req.Password != "" {
		params.Password = req.Password
	}

	a, err := s.accountService.Update(ctx, params)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return a.ToGRPCAccount(), nil
}
