package repository

import (
	"context"

	"github.com/clubo-app/clubben/search-service/datastruct"
	"github.com/jonashiltl/govespa"
)

const (
	DOC_TYPE  = "user"
	NAMESPACE = "default"
)

type ProfileRepository struct {
	v *govespa.VespaClient
}

func NewProfileRepository(v *govespa.VespaClient) ProfileRepository {
	return ProfileRepository{
		v: v,
	}
}

func (r *ProfileRepository) QueryProfile(c context.Context, query string) ([]datastruct.Profile, error) {
	profiles := make([]datastruct.Profile, 0, 20)

	_, err := r.v.
		Query().
		WithContext(c).
		AddYQL("select * from user where userInput(@q)").
		AddVariable("q", query).
		AddParameter(govespa.QueryParameter{
			Hits: 20,
		}).
		Select(&profiles)
	if err != nil {
		return []datastruct.Profile{}, err
	}

	return profiles, nil
}

func (r *ProfileRepository) PutProfile(c context.Context, u datastruct.Profile) error {
	return r.v.
		Put(govespa.DocumentId{Namespace: NAMESPACE, DocType: DOC_TYPE, UserSpecific: u.Id}).
		WithContext(c).
		BindStruct(u).
		Exec()
}

func (r *ProfileRepository) UpdateProfile(c context.Context, u datastruct.Profile) error {
	update := r.v.
		Update(govespa.DocumentId{Namespace: NAMESPACE, DocType: DOC_TYPE, UserSpecific: u.Id}).
		WithContext(c)

	if u.Username != "" {
		update.Assign("username", u.Username)
	}
	if u.Firstname != "" {
		update.Assign("firstname", u.Firstname)
	}
	if u.Lastname != "" {
		update.Assign("lastname", u.Lastname)
	}

	return update.Exec()
}

func (r *ProfileRepository) RemoveProfile(c context.Context, id string) error {
	return r.v.
		Remove(govespa.DocumentId{Namespace: NAMESPACE, DocType: DOC_TYPE, UserSpecific: id}).
		WithContext(c).
		Exec()
}

func (r *ProfileRepository) IncrementFriendCount(c context.Context, id string) error {
	return r.v.
		Update(govespa.DocumentId{Namespace: NAMESPACE, DocType: DOC_TYPE, UserSpecific: id}).
		Increment("friend_count", 1).
		Exec()
}

func (r *ProfileRepository) DecrementFriendCount(c context.Context, id string) error {
	return r.v.
		Update(govespa.DocumentId{Namespace: NAMESPACE, DocType: DOC_TYPE, UserSpecific: id}).
		Decrement("friend_count", 1).
		Exec()
}
