package rpc

import (
	"context"

	pg "github.com/clubo-app/protobuf/profile"
)

func (s *profileServer) UsernameTaken(ctx context.Context, req *pg.UsernameTakenRequest) (*pg.UsernameTakenResponse, error) {
	usernameTaken := s.ps.UsernameTaken(ctx, req.Username)

	return &pg.UsernameTakenResponse{Taken: usernameTaken}, nil
}
