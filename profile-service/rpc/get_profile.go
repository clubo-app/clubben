package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	pg "github.com/clubo-app/clubben/protobuf/profile"
)

func (s *profileServer) GetProfile(ctx context.Context, req *pg.GetProfileRequest) (*pg.Profile, error) {
	p, err := s.ps.GetById(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCProfile(), nil
}
