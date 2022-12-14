package service

import (
	"context"

	"github.com/clubo-app/clubben/relation-service/datastruct"
)

type FriendRelation interface {
	CreateFriendRequest(ctx context.Context, uId, fId string) (datastruct.FriendRelation, error)
	DeclineFriendRequest(ctx context.Context, uId, fId string) error
	AcceptFriendRequest(ctx context.Context, uId, fId string) error
	RemoveFriendRelation(ctx context.Context, uId, fId string) error
	GetFriendRelation(ctx context.Context, uId, fId string) (datastruct.FriendRelation, error)
	GetFriends(ctx context.Context, uId string, page []byte, limit uint64) ([]datastruct.FriendRelation, []byte, error)
	GetIncomingFriendRequests(ctx context.Context, uId string, page []byte, limit uint64) ([]datastruct.FriendRelation, []byte, error)
	IncreaseFriendCount(ctx context.Context, uId string) error
	DecreaseFriendCount(ctx context.Context, uId string) error
	GetFriendCount(ctx context.Context, uId string) (datastruct.FriendCount, error)
	GetManyFriendCount(ctx context.Context, ids []string) ([]datastruct.FriendCount, error)
}
