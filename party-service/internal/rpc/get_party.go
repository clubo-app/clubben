package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s partyServer) GetParty(c context.Context, req *pbparty.GetPartyRequest) (*pbparty.Party, error) {
	id, err := ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	p, err := s.ps.Get(c, id.String())
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCParty(), nil
}
