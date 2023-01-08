package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
)

func (s partyServer) GetManyPartiesMap(ctx context.Context, req *pbparty.GetManyPartiesRequest) (*pbparty.GetManyPartiesMapResponse, error) {
	ps, err := s.ps.GetMany(ctx, req.Ids)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	parties := make(map[string]*pbparty.Party)

	for _, p := range ps {
		parties[p.ID] = p.ToGRPCParty()
	}

	return &pbparty.GetManyPartiesMapResponse{Parties: parties}, nil
}
