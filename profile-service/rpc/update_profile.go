package rpc

import (
	"context"
	"strings"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/profile-service/dto"
	pg "github.com/clubo-app/clubben/protobuf/profile"
)

func (s *profileServer) UpdateProfile(ctx context.Context, req *pg.UpdateProfileRequest) (*pg.Profile, error) {
	dp := dto.Profile{
		ID:        req.Id,
		Username:  strings.ToLower(req.Username),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Avatar:    req.Avatar,
	}

	if dp.Avatar != "" {
		loc, err := s.up.Upload(ctx, req.Id, dp.Avatar)
		if err != nil {
			return nil, utils.HandleError(err)
		}
		dp.Avatar = loc
	}

	p, err := s.ps.Update(ctx, dp)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCProfile(), nil
}
