package rpc

import (
	"context"

	"github.com/clubo-app/packages/utils"
	pg "github.com/clubo-app/protobuf/profile"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *profileServer) GetProfile(ctx context.Context, req *pg.GetProfileRequest) (*pg.Profile, error) {
	_, err := ksuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}

	p, err := s.ps.GetById(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCProfile(), nil
}
