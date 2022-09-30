package repository

import (
	"context"
	"errors"

	"github.com/clubo-app/clubben/relation-service/datastruct"
	"github.com/go-playground/validator/v10"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

const (
	FAVORITE_PARTIES         string = "favorite_parties"
	FAVORITE_PARTIES_BY_USER string = "favorite_parties_by_user"
	FAVORITE_PARTY_COUNT     string = "favorite_party_count"
)

var favoritePartyMetadata = table.Metadata{
	Name:    FAVORITE_PARTIES,
	Columns: []string{"user_id", "party_id", "favorited_at"},
	PartKey: []string{"user_id", "party_id"},
}

type FavoritePartyRepository interface {
	FavorParty(ctx context.Context, fp datastruct.FavoriteParty) (datastruct.FavoriteParty, error)
	DefavorParty(ctx context.Context, uId, pId string) error
	GetFavoriteParty(ctx context.Context, uId, pId string) (datastruct.FavoriteParty, error)
	GetFavoritePartyManyUser(ctx context.Context, uId []string, pId string) ([]datastruct.FavoriteParty, error)
	GetFavoritePartiesByUser(ctx context.Context, uId string, page []byte, limit uint64) ([]datastruct.FavoriteParty, []byte, error)
	GetFavorisingUsersByParty(ctx context.Context, pId string, page []byte, limit uint64) ([]datastruct.FavoriteParty, []byte, error)
}

type favoritePartyRepository struct {
	sess *gocqlx.Session
	val  *validator.Validate
}

func (r *favoritePartyRepository) FavorParty(ctx context.Context, fp datastruct.FavoriteParty) (datastruct.FavoriteParty, error) {
	err := r.val.Struct(fp)
	if err != nil {
		return datastruct.FavoriteParty{}, err
	}

	stmt, names := qb.
		Insert(FAVORITE_PARTIES).
		Columns(favoritePartyMetadata.Columns...).
		Unique().
		ToCql()

	err = r.sess.
		Query(stmt, names).
		BindStruct(fp).
		ExecRelease()
	if err != nil {
		return datastruct.FavoriteParty{}, err
	}

	return fp, nil
}

func (r *favoritePartyRepository) DefavorParty(ctx context.Context, uId, pId string) error {
	stmt, names := qb.
		Delete(FAVORITE_PARTIES).
		Where(qb.Eq("user_id")).
		Where(qb.Eq("party_id")).
		Existing().
		ToCql()

	err := r.sess.
		Query(stmt, names).
		BindMap((qb.M{"party_id": pId, "user_id": uId})).
		ExecRelease()
	if err != nil {
		return err
	}

	return nil
}

func (r *favoritePartyRepository) GetFavoriteParty(ctx context.Context, uId, pId string) (datastruct.FavoriteParty, error) {
	var res datastruct.FavoriteParty
	stmt, names := qb.
		Select(FAVORITE_PARTIES).
		Where(qb.Eq("user_id")).
		Where(qb.Eq("party_id")).
		ToCql()

	err := r.sess.
		Query(stmt, names).
		BindMap((qb.M{"party_id": pId, "user_id": uId})).
		GetRelease(&res)
	if err != nil {
		return datastruct.FavoriteParty{}, err
	}

	return res, nil
}

func (r *favoritePartyRepository) GetFavoritePartyManyUser(ctx context.Context, uId []string, pId string) ([]datastruct.FavoriteParty, error) {
	res := make([]datastruct.FavoriteParty, len(uId))

	stmt, names := qb.
		Select(FAVORITE_PARTIES).
		Where(qb.In("user_id")).
		Where(qb.Eq("party_id")).
		ToCql()

	err := r.sess.
		Query(stmt, names).
		BindMap((qb.M{
			"party_id": pId,
			"user_id":  uId,
		})).
		SelectRelease(&res)
	if err != nil {
		return make([]datastruct.FavoriteParty, 0), err
	}

	return res, nil
}

func (r *favoritePartyRepository) GetFavoritePartiesByUser(ctx context.Context, uId string, page []byte, limit uint64) (result []datastruct.FavoriteParty, nextPage []byte, err error) {
	stmt, names := qb.
		Select(FAVORITE_PARTIES_BY_USER).
		Where(qb.Eq("user_id")).
		ToCql()

	q := r.sess.
		Query(stmt, names).
		BindMap((qb.M{"user_id": uId}))
	defer q.Release()

	q.PageState(page)
	if limit == 0 {
		q.PageSize(10)
	} else {
		q.PageSize(int(limit))
	}

	iter := q.Iter()
	err = iter.Select(&result)
	if err != nil {
		return []datastruct.FavoriteParty{}, nil, errors.New("no favorite parties found")
	}

	return result, iter.PageState(), nil
}

func (r *favoritePartyRepository) GetFavorisingUsersByParty(ctx context.Context, pId string, page []byte, limit uint64) (result []datastruct.FavoriteParty, nextPage []byte, err error) {
	stmt, names := qb.
		Select(FAVORITE_PARTIES).
		Where(qb.Eq("party_id")).
		ToCql()

	q := r.sess.
		Query(stmt, names).
		BindMap((qb.M{"party_id": pId}))
	defer q.Release()

	q.PageState(page)
	if limit == 0 {
		q.PageSize(10)
	} else {
		q.PageSize(int(limit))
	}

	iter := q.Iter()
	err = iter.Select(&result)
	if err != nil {
		return []datastruct.FavoriteParty{}, nil, errors.New("no favorite parties found")
	}

	return result, iter.PageState(), nil
}
