package repository

import (
	"context"
	"log"

	"github.com/clubo-app/clubben/search-service/datastruct"
	"github.com/jonashiltl/govespa"
)

type ProfileRepository struct {
	v *govespa.VespaClient
}

func NewProfileRepository(v *govespa.VespaClient) ProfileRepository {
	return ProfileRepository{
		v: v,
	}
}

func (r *ProfileRepository) QueryProfile(c context.Context, query, continuation string) ([]datastruct.Profile, string, error) {
	profiles := new(datastruct.Profile)

	res, err := r.v.
		Query().
		WithContext(c).
		AddYQL("select * from user where userInput(@q)").
		AddVariable("q", query).
		AddParameter(govespa.QueryParameter{}).
		Get(&profiles)
	log.Printf("QUERY RESULT: %+v", res)
	if err != nil {
		return []datastruct.Profile{}, "", err
	}

	return []datastruct.Profile{*profiles}, "", nil
}

func (r *ProfileRepository) PutProfile(c context.Context, u datastruct.Profile) error {
	return r.v.
		Put(govespa.DocumentId{Namespace: "default", DocType: "user", UserSpecific: u.Id}).
		WithContext(c).
		BindStruct(u).
		Exec()
}

func (r *ProfileRepository) UpdateProfile(c context.Context, u datastruct.Profile) error {
	update := r.v.
		Update(govespa.DocumentId{Namespace: "default", DocType: "user", UserSpecific: u.Id}).
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
		Remove(govespa.DocumentId{Namespace: "default", DocType: "user", UserSpecific: id}).
		WithContext(c).
		Exec()
}
