package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	ag "github.com/clubo-app/clubben/protobuf/auth"
	cg "github.com/clubo-app/clubben/protobuf/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) DeleteAccount(ctx context.Context, req *ag.DeleteAccountRequest) (*cg.SuccessIndicator, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "empty user id")
	}

	err := s.ac.Delete(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &cg.SuccessIndicator{Sucess: true}, nil
}
