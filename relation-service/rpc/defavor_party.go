package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	cg "github.com/clubo-app/clubben/protobuf/common"
	rg "github.com/clubo-app/clubben/protobuf/relation"
)

func (s *relationServer) DefavorParty(ctx context.Context, req *rg.PartyAndUserRequest) (*cg.SuccessIndicator, error) {
	err := s.fp.DefavorParty(ctx, req.UserId, req.PartyId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &cg.SuccessIndicator{Sucess: true}, nil
}
