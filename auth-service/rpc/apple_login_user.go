package rpc

import (
	"context"

	ag "github.com/clubo-app/clubben/protobuf/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) AppleLoginUser(ctx context.Context, req *ag.AppleLoginUserRequest) (*ag.LoginUserResponse, error) {
	return nil, status.Error(codes.Internal, "Not yet implemented")
}
