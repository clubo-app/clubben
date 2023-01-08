package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/party-service/internal/repository"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
)

func (s partyServer) GetByUser(c context.Context, req *pbparty.GetByUserRequest) (*pbparty.PagedParties, error) {
	ps, err := s.ps.GetByUser(c, repository.GetPartiesByUserParams{
		UserID:   req.UserId,
		IsPublic: req.IsPublic,
		Limit:    req.Limit,
		Offset:   req.Offset,
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	var pp []*pbparty.Party
	for _, p := range ps {
		pp = append(pp, p.ToGRPCParty())
	}

	return &pbparty.PagedParties{Parties: pp}, nil
}
