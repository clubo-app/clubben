package rpc

import (
	"context"

	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
)

func (s *authServer) EmailTaken(ctx context.Context, req *pbauth.EmailTakenRequest) (*pbauth.EmailTakenResponse, error) {
	taken := s.accountService.EmailTaken(ctx, req.Email)

	return &pbauth.EmailTakenResponse{Taken: taken}, nil
}
