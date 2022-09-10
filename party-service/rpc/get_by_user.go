package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/party-service/repository"
	pg "github.com/clubo-app/clubben/protobuf/party"
)

func (s partyServer) GetByUser(c context.Context, req *pg.GetByUserRequest) (*pg.PagedParties, error) {
	ps, err := s.ps.GetByUser(c, repository.GetPartiesByUserParams{
		UserID:   req.UserId,
		IsPublic: req.IsPublic,
		Limit:    req.Limit,
		Offset:   req.Offset,
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	var pp []*pg.Party
	for _, p := range ps {
		pp = append(pp, p.ToGRPCParty())
	}

	return &pg.PagedParties{Parties: pp}, nil
}
