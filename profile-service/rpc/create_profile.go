package rpc

import (
	"context"
	"strings"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/profile-service/dto"
	pg "github.com/clubo-app/protobuf/profile"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *profileServer) CreateProfile(ctx context.Context, req *pg.CreateProfileRequest) (*pg.Profile, error) {
	if strings.Contains(req.Username, "@") {
		return nil, status.Error(codes.InvalidArgument, "Not allowed to have @ in your username")
	}

	p, err := s.ps.Create(ctx, dto.Profile{
		ID:        req.Id,
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Avatar:    req.Avatar,
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCProfile(), nil
}
