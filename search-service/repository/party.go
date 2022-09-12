package repository

import (
	"context"

	"github.com/clubo-app/clubben/search-service/datastruct"
	"github.com/jonashiltl/govespa"
)

const (
	PARTY_DOC_TYPE  = "party"
	PARTY_NAMESPACE = "default"
)

type PartyRepository struct {
	v *govespa.VespaClient
}

func NewPartyRepository(v *govespa.VespaClient) PartyRepository {
	return PartyRepository{
		v: v,
	}
}

func (r *PartyRepository) QueryParty(c context.Context, query string, loc datastruct.Location) ([]datastruct.Party, error) {
	parties := make([]datastruct.Party, 0, 15)

	_, err := r.v.
		Query().
		WithContext(c).
		AddYQL("select * from party where userInput(@q)").
		AddVariable("q", query).
		AddParameter(govespa.QueryParameter{
			Hits: 15,
		}).
		Select(&parties)
	if err != nil {
		return []datastruct.Party{}, err
	}

	return parties, nil

}

func (r *PartyRepository) PutParty(c context.Context, p datastruct.Party) error {
	return r.v.
		Put(govespa.DocumentId{Namespace: NAMESPACE, DocType: PARTY_DOC_TYPE, UserSpecific: p.Id}).
		WithContext(c).
		BindStruct(p).
		Exec()
}

func (r *PartyRepository) UpdateParty(c context.Context, p datastruct.Party) error {
	update := r.v.
		Update(govespa.DocumentId{Namespace: NAMESPACE, DocType: PARTY_DOC_TYPE, UserSpecific: p.Id}).
		WithContext(c)

	if p.Title != "" {
		update.Assign("title", p.Title)
	}
	if p.Description != "" {
		update.Assign("description", p.Description)
	}
	if p.MusicGenre != "" {
		update.Assign("music_genre", p.MusicGenre)
	}
	if p.EntryDate != 0 {
		update.Assign("entry_date", p.EntryDate)
	}
	if p.FavoriteCount != 0 {
		update.Assign("favorite_count", p.FavoriteCount)
	}
	if p.Location.Lat != 0 && p.Location.Lng != 0 {
		update.Assign("location", p.Location)
	}

	return update.Exec()
}

func (r *PartyRepository) RemoveParty(c context.Context, id string) error {
	return r.v.
		Remove(govespa.DocumentId{Namespace: NAMESPACE, DocType: PARTY_DOC_TYPE, UserSpecific: id}).
		WithContext(c).
		Exec()
}

func (r *PartyRepository) IncrementFavoriteCount(c context.Context, id string) error {
	return r.v.
		Update(govespa.DocumentId{Namespace: NAMESPACE, DocType: PARTY_DOC_TYPE, UserSpecific: id}).
		Increment("favorite_count", 1).
		Exec()
}

func (r *PartyRepository) DecrementFavoriteCount(c context.Context, id string) error {
	return r.v.
		Update(govespa.DocumentId{Namespace: NAMESPACE, DocType: PARTY_DOC_TYPE, UserSpecific: id}).
		Decrement("favorite_count", 1).
		Exec()
}
