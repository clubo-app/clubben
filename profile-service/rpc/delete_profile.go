package rpc

import (
	"context"

	"github.com/clubo-app/packages/utils"
	cg "github.com/clubo-app/protobuf/common"
	pg "github.com/clubo-app/protobuf/profile"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *profileServer) DeleteProfile(ctx context.Context, req *pg.DeleteProfileRequest) (*cg.SuccessIndicator, error) {
	_, err := ksuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}

	err = s.ps.Delete(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &cg.SuccessIndicator{Sucess: true}, nil
}
