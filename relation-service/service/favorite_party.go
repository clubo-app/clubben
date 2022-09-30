package service

import (
	"context"

	"github.com/clubo-app/clubben/relation-service/datastruct"
)

type FavoriteParty interface {
	FavorParty(ctx context.Context, fp datastruct.FavoriteParty) (datastruct.FavoriteParty, error)
	DefavorParty(ctx context.Context, uId, pId string) error
	GetFavoriteParty(ctx context.Context, uId, pId string) (datastruct.FavoriteParty, error)
	GetFavoritePartyManyUser(ctx context.Context, uId []string, pId string) ([]datastruct.FavoriteParty, error)
	GetFavoritePartiesByUser(ctx context.Context, uId string, page []byte, limit uint64) ([]datastruct.FavoriteParty, []byte, error)
	GetFavorisingUsersByParty(ctx context.Context, pId string, page []byte, limit uint64) ([]datastruct.FavoriteParty, []byte, error)
}
