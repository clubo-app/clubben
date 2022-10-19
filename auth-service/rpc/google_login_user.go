package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	ag "github.com/clubo-app/clubben/protobuf/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) GoogleLoginUser(ctx context.Context, req *ag.GoogleLoginUserRequest) (*ag.LoginUserResponse, error) {
	googleClaims, err := s.goog.ValidateGoogleToken(ctx, req.Token)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Google Token")
	}

	account, err := s.accountService.GetByEmail(ctx, googleClaims.Email)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	at, err := s.token.NewAccessToken(account)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	rt, err := s.token.NewRefreshToken(account.ID, account.RefreshTokenGeneration)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &ag.LoginUserResponse{
		Account: account.ToGRPCAccount(),
		Tokens: &ag.TokenResponse{
			AccessToken:  at,
			RefreshToken: rt,
		},
	}, nil
}
