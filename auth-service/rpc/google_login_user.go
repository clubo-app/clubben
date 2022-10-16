package rpc

import (
	"context"

	ag "github.com/clubo-app/clubben/protobuf/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) GoogleLoginUser(ctx context.Context, req *ag.GoogleLoginUserRequest) (*ag.LoginUserResponse, error) {
	return nil, status.Error(codes.Unavailable, "not yet implemented")
}
