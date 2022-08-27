package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	pg "github.com/clubo-app/clubben/protobuf/party"
)

func (s partyServer) GetManyPartiesMap(ctx context.Context, req *pg.GetManyPartiesRequest) (*pg.GetManyPartiesMapResponse, error) {
	ps, err := s.ps.GetMany(ctx, req.Ids)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	parties := make(map[string]*pg.Party)

	for _, p := range ps {
		parties[p.ID] = p.ToGRPCParty()
	}

	return &pg.GetManyPartiesMapResponse{Parties: parties}, nil
}
