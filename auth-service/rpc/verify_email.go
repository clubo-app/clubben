package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	ag "github.com/clubo-app/clubben/protobuf/auth"
)

func (s *authServer) VerifyEmail(ctx context.Context, req *ag.VerifyEmailRequest) (*ag.Account, error) {
	a, err := s.accountService.UpdateVerified(ctx, req.Id, req.Code, true)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return a.ToGRPCAccount(), nil
}
