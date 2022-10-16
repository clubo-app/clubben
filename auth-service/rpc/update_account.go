package rpc

import (
	"context"
	"errors"
	"net/mail"

	"github.com/clubo-app/clubben/auth-service/repository"
	"github.com/clubo-app/clubben/libs/utils"
	ag "github.com/clubo-app/clubben/protobuf/auth"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) UpdateAccount(ctx context.Context, req *ag.UpdateAccountRequest) (*ag.Account, error) {
	id, err := ksuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	params := repository.UpdateAccountParams{
		ID: id.String(),
	}

	if req.Password != "" {
		hash, err := s.pw.HashPassword(req.Password)
		if err != nil {
			return nil, utils.HandleError(err)
		}
		params.PasswordHash = hash
	}

	if req.Email != "" {
		_, err = mail.ParseAddress(req.Email)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Invalid Email")
		}
		params.Email = req.Email
	}

	a, err := s.ac.Update(ctx, params)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return a.ToGRPCAccount(), nil
}
