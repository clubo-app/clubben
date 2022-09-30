package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/search"
)

func (s *searchServer) SearchUsers(ctx context.Context, req *search.SearchUsersRequest) (*search.SearchUsersResponse, error) {
	profiles, err := s.profile.QueryProfile(ctx, req.Query)
	if err != nil {
		return &search.SearchUsersResponse{}, utils.HandleError(err)
	}

	res := make([]*search.IndexedUser, len(profiles))
	for i, p := range profiles {
		res[i] = p.ToGRPCUser()
	}

	return &search.SearchUsersResponse{
		Users: res,
	}, nil
}
