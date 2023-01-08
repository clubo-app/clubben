package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
)

func (s partyServer) GetManyParties(ctx context.Context, req *pbparty.GetManyPartiesRequest) (*pbparty.GetManyPartiesResponse, error) {
	ps, err := s.ps.GetMany(ctx, req.Ids)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	parties := make([]*pbparty.Party, len(ps))

	for i, p := range ps {
		parties[i] = p.ToGRPCParty()
	}

	return &pbparty.GetManyPartiesResponse{Parties: parties}, nil
}
