package rpc

import (
	"context"
	"errors"
	"strings"

	"github.com/clubo-app/clubben/profile-service/dto"
	"github.com/clubo-app/packages/utils"
	pg "github.com/clubo-app/protobuf/profile"
	"github.com/segmentio/ksuid"
)

func (s *profileServer) UpdateProfile(ctx context.Context, req *pg.UpdateProfileRequest) (*pg.Profile, error) {
	id, err := ksuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	dp := dto.Profile{
		ID:        id.String(),
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
