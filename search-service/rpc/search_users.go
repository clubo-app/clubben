package rpc

import (
	"context"

	"github.com/clubo-app/clubben/protobuf/search"
	"github.com/clubo-app/packages/utils"
)

func (s *searchServer) SearchUsers(ctx context.Context, req *search.SearchUsersRequest) (*search.PagedIndexedUsers, error) {
	profiles, _, err := s.profile.QueryProfile(ctx, req.Query, "")
	if err != nil {
		return &search.PagedIndexedUsers{}, utils.HandleError(err)
	}

	res := make([]*search.IndexedUser, len(profiles))
	for i, p := range profiles {
		res[i] = p.ToGRPCProfile()
	}

	return &search.PagedIndexedUsers{
		Users: res,
	}, nil
}
