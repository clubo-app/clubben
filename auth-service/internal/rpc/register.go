package rpc

import (
	"context"
	"net/mail"

	"github.com/clubo-app/clubben/auth-service/internal/repository"
	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	"github.com/clubo-app/clubben/libs/id"
	"github.com/clubo-app/clubben/libs/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) Register(ctx context.Context, req *pbauth.RegisterRequest) (*pbauth.RegisterResponse, error) {
	params := repository.CreateAccountParams{
		ID:       id.New(id.User),
		Email:    req.Email,
		Password: req.Password,
	}

	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Email")
	}

	a, err := s.accountService.Create(ctx, params)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &pbauth.RegisterResponse{Account: a.ToGRPCAccount()}, nil
}
