package service

import (
	"context"

	"github.com/clubo-app/clubben/story-service/datastruct"
	"github.com/clubo-app/clubben/story-service/dto"
	"github.com/clubo-app/clubben/story-service/repository"
)

const GEOHASH_PRECISION uint = 9

type StoryService interface {
	Create(context.Context, dto.Story) (datastruct.Story, error)
	Delete(context.Context, repository.DeleteParams) error
	GetByUser(context.Context, repository.GetByUserParams) ([]datastruct.Story, []byte, error)
	GetPastByUser(context.Context, repository.GetByUserParams) ([]datastruct.Story, []byte, error)
	GetByParty(context.Context, repository.GetByPartyParams) ([]datastruct.Story, []byte, error)
}
