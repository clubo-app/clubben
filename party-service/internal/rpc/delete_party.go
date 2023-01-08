package rpc

import (
	"context"

	"github.com/clubo-app/clubben/libs/utils"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s partyServer) DeleteParty(c context.Context, req *pbparty.DeletePartyRequest) (*emptypb.Empty, error) {
	_, err := ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	_, err = ksuid.Parse(req.RequesterId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Requester id")
	}

	err = s.ps.Delete(c, req.RequesterId, req.PartyId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &emptypb.Empty{}, nil
}
