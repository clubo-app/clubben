package rpc

import (
	"context"

	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	"github.com/clubo-app/clubben/libs/utils"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *authServer) DeleteAccount(ctx context.Context, req *pbauth.DeleteAccountRequest) (*emptypb.Empty, error) {
	err := s.accountService.Delete(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &emptypb.Empty{}, nil
}
