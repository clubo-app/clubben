package rpc

import (
	"context"
	"database/sql"
	"net/mail"

	"github.com/clubo-app/clubben/auth-service/repository"
	"github.com/clubo-app/clubben/libs/utils"
	ag "github.com/clubo-app/clubben/protobuf/auth"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) RegisterUser(ctx context.Context, req *ag.RegisterUserRequest) (*ag.LoginUserResponse, error) {
	code, _ := utils.GenerateOTP(4)

	if req.GoogleToken != "" && req.AppleToken != "" {
		return nil, status.Error(codes.InvalidArgument, "Can't use both Google & Apple Token")
	}

	params := repository.CreateAccountParams{
		ID:            ksuid.New().String(),
		Email:         req.Email,
		EmailVerified: false,
		EmailCode:     sql.NullString{Valid: code != "", String: code},
		Type:          repository.TypeUSER,
	}

	if req.Password != "" {
		hash, err := s.pw.HashPassword(req.Password)
		if err != nil {
			return nil, utils.HandleError(err)
		}
		params.PasswordHash = sql.NullString{Valid: hash != "", String: hash}
	}

	if req.GoogleToken != "" {
		googleClaims, err := s.goog.ValidateGoogleToken(ctx, req.GoogleToken)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Invalid Google Token")
		}

		params.Email = googleClaims.Email
		params.EmailVerified = googleClaims.EmailVerified
		params.Provider = repository.NullProvider{Valid: true, Provider: repository.ProviderGOOGLE}
		params.PasswordHash = sql.NullString{Valid: false}
	}

	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Email")
	}

	a, err := s.accountService.Create(ctx, params)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	at, err := s.token.NewAccessToken(a)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	rt, err := s.token.NewRefreshToken(a.ID, a.RefreshTokenGeneration)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &ag.LoginUserResponse{
		Tokens: &ag.TokenResponse{
			AccessToken:  at,
			RefreshToken: rt,
		},
		Account: a.ToGRPCAccount(),
	}, nil
}
