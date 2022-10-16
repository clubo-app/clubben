package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	ag "github.com/clubo-app/clubben/protobuf/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *authServer) DeleteAccount(ctx context.Context, req *ag.DeleteAccountRequest) (*emptypb.Empty, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "empty user id")
	}

	err := s.ac.Delete(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &emptypb.Empty{}, nil
}
