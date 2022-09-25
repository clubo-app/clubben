// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package repository

import (
	"context"
)

type Querier interface {
	DecreaseFavoriteCount(ctx context.Context, arg DecreaseFavoriteCountParams) error
	DecreaseParticipantsCount(ctx context.Context, arg DecreaseParticipantsCountParams) error
	DeleteParty(ctx context.Context, arg DeletePartyParams) error
	IncreaseFavoriteCount(ctx context.Context, arg IncreaseFavoriteCountParams) error
	IncreaseParticipantsCount(ctx context.Context, arg IncreaseParticipantsCountParams) error
}

var _ Querier = (*Queries)(nil)
