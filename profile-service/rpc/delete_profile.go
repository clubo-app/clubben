package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	cg "github.com/clubo-app/clubben/protobuf/common"
	pg "github.com/clubo-app/clubben/protobuf/profile"
)

func (s *profileServer) DeleteProfile(ctx context.Context, req *pg.DeleteProfileRequest) (*cg.SuccessIndicator, error) {
	err := s.ps.Delete(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &cg.SuccessIndicator{Sucess: true}, nil
}
