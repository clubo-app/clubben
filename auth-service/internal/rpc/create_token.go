package rpc

import (
	"context"

	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	"github.com/clubo-app/clubben/libs/utils"
)

func (s authServer) CreateToken(ctx context.Context, req *pbauth.CreateTokenRequest) (*pbauth.CreateTokenResponse, error) {
	token, err := s.accountService.CreateToken(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}
	return &pbauth.CreateTokenResponse{Token: token}, nil
}
