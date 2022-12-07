package rpc

import (
	"context"

	ag "github.com/clubo-app/clubben/protobuf/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *authServer) ResendVerificationEmail(context.Context, *ag.ResendVerificationEmailRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unavailable, "not yet implemented")

}
