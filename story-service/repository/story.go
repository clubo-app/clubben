package repository

import (
	"context"
	"errors"
	"time"

	"github.com/clubo-app/clubben/story-service/datastruct"
	"github.com/clubo-app/clubben/story-service/dto"
	"github.com/go-playground/validator/v10"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	STORIES_BY_USER  string = "stories_by_user"
	STORIES_BY_PARTY string = "stories_by_party"
)

type StoryRepository interface {
	Create(context.Context, dto.Story) (datastruct.Story, error)
	Delete(context.Context, DeleteParams) error
	GetByUser(context.Context, GetByUserParams) ([]datastruct.Story, []byte, error)
	GetPastByUser(context.Context, GetByUserParams) ([]datastruct.Story, []byte, error)
	GetByParty(context.Context, GetByPartyParams) ([]datastruct.Story, []byte, error)
}

var storyMetadata = table.Metadata{
	Name:    STORIES_BY_USER,
	Columns: []string{"story_id", "party_id", "user_id", "url", "tagged_friends"},
	PartKey: []string{"user_id", "story_id"},
}

type storyRepository struct {
	sess *gocqlx.Session
	val  *validator.Validate
}

func (r storyRepository) Create(ctx context.Context, ds dto.Story) (datastruct.Story, error) {
	s := datastruct.Story{
		Id:            ds.Id.String(),
		PartyId:       ds.PartyId,
		UserId:        ds.UserId,
		Url:           ds.Url,
		TaggedFriends: ds.TaggedFriends,
	}

	err := r.val.StructCtx(ctx, s)
	if err != nil {
		return datastruct.Story{}, err
	}

	stmt, names := qb.
		Insert(STORIES_BY_USER).
		Columns(storyMetadata.Columns...).
		TTL(time.Hour * 24).
		ToCql()

	err = r.sess.
		ContextQuery(ctx, stmt, names).
		BindStruct(s).
		ExecRelease()
	if err != nil {
		return datastruct.Story{}, err
	}

	return s, nil
}

type DeleteParams struct {
	SId string
	UId string
}

func (r storyRepository) Delete(ctx context.Context, params DeleteParams) error {
	stmt, names := qb.
		Delete(STORIES_BY_USER).
		Where(qb.Eq("story_id")).
		Where(qb.Eq("user_id")).
		ToCql()

	err := r.sess.
		ContextQuery(ctx, stmt, names).
		BindMap((qb.M{
			"story_id": params.SId,
			"user_id":  params.UId,
		})).
		ExecRelease()
	if err != nil {
		return err
	}

	return nil
}

type GetByUserParams struct {
	UId   string
	Limit int
	Page  []byte
}

func (r storyRepository) GetByUser(ctx context.Context, params GetByUserParams) (res []datastruct.Story, page []byte, err error) {
	stmt, names := qb.
		Select(STORIES_BY_USER).
		Where(qb.Eq("user_id")).
		Where(qb.GtOrEq("story_id")).
		ToCql()

	comp, err := ksuid.NewRandomWithTime(time.Now().AddDate(0, 0, -1))
	if err != nil {
		return []datastruct.Story{}, nil, status.Error(codes.Internal, "Failed to generate yesterdays timestamp")
	}

	q := r.sess.
		ContextQuery(ctx, stmt, names).
		BindMap((qb.M{
			"user_id":  params.UId,
			"story_id": comp.String(),
		}))
	defer q.Release()

	q.PageState(params.Page)
	if params.Limit == 0 {
		q.PageSize(20)
	} else {
		q.PageSize(params.Limit)
	}

	iter := q.Iter()
	err = iter.Select(&res)
	if err != nil {
		return []datastruct.Story{}, nil, errors.New("no stories found")
	}

	return res, iter.PageState(), nil
}

type GetByPartyParams struct {
	PId   string
	Limit int
	Page  []byte
}

func (r storyRepository) GetByParty(ctx context.Context, params GetByPartyParams) (res []datastruct.Story, page []byte, err error) {
	stmt, names := qb.
		Select(STORIES_BY_PARTY).
		Where(qb.Eq("party_id")).
		ToCql()

	q := r.sess.
		ContextQuery(ctx, stmt, names).
		BindMap((qb.M{
			"party_id": params.PId,
		}))
	defer q.Release()

	q.PageState(params.Page)
	if params.Limit == 0 {
		q.PageSize(20)
	} else {
		q.PageSize(params.Limit)
	}

	iter := q.Iter()
	err = iter.Select(&res)
	if err != nil {
		return []datastruct.Story{}, nil, errors.New("no stories found")
	}

	return res, iter.PageState(), nil
}

func (r storyRepository) GetPastByUser(ctx context.Context, params GetByUserParams) (res []datastruct.Story, page []byte, err error) {
	stmt, names := qb.
		Select(STORIES_BY_USER).
		Where(qb.Eq("user_id")).
		Where(qb.Lt("story_id")).
		ToCql()

	comp, err := ksuid.NewRandomWithTime(time.Now().AddDate(0, 0, -1))
	if err != nil {
		return []datastruct.Story{}, nil, status.Error(codes.Internal, "Failed to generate yesterdays timestamp")
	}

	q := r.sess.
		ContextQuery(ctx, stmt, names).
		BindMap((qb.M{
			"user_id":  params.UId,
			"story_id": comp.String(),
		}))
	defer q.Release()

	q.PageState(params.Page)
	if params.Limit == 0 {
		q.PageSize(20)
	} else {
		q.PageSize(params.Limit)
	}

	iter := q.Iter()
	err = iter.Select(&res)
	if err != nil {
		return []datastruct.Story{}, nil, errors.New("no stories found")
	}

	return res, iter.PageState(), nil
}
