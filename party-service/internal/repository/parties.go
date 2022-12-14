package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/leporo/sqlf"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/ewkb"
	"github.com/paulmach/orb/encoding/wkb"
)

const (
	TABLE_NAME = "parties"
)

type PartyRepository struct {
	Pool    *pgxpool.Pool
	querier Querier
}

func NewPartyRepository(urlStr string) (*PartyRepository, error) {
	log.Println("Conntected to ", urlStr)
	pgURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	connURL := *pgURL
	if connURL.Scheme == "cockroachdb" {
		connURL.Scheme = "postgres"
	}

	c, err := pgxpool.ParseConfig(connURL.String())
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), c)
	if err != nil {
		return nil, fmt.Errorf("pgx connection error: %w", err)
	}

	err = validateSchema(connURL)
	if err != nil {
		log.Printf("Schema validation error: %v", err)
	}

	return &PartyRepository{
		Pool:    pool,
		querier: New(pool),
	}, nil
}

func (r PartyRepository) Close() {
	r.Pool.Close()
}

type CreatePartyParams struct {
	ID              string
	UserID          string
	Title           string
	Description     sql.NullString
	IsPublic        bool
	MusicGenre      string
	MaxParticipants int32
	Location        orb.Point
	StreetAddress   string
	PostalCode      string
	State           string
	Country         string
	EntryDate       time.Time
}

const selectStmt = "id, user_id, title, description, is_public, music_genre, ST_AsBinary(location) AS location, street_address, postal_code, state, country, entry_date, max_participants, participants_count, favorite_count"

func (r PartyRepository) CreateParty(ctx context.Context, arg CreatePartyParams) (Party, error) {
	sqlf.SetDialect(sqlf.PostgreSQL)
	b := sqlf.
		InsertInto(TABLE_NAME).
		Set("id", arg.ID).
		Set("user_id", arg.UserID).
		Set("title", arg.Title).
		Set("description", arg.Description).
		Set("is_public", arg.IsPublic).
		Set("music_genre", arg.MusicGenre).
		Set("max_participants", arg.MaxParticipants).
		SetExpr("location", "ST_GeomFromEWKB(?)", ewkb.Value(arg.Location, 4326)).
		Set("street_address", arg.StreetAddress).
		Set("postal_code", arg.PostalCode).
		Set("state", arg.State).
		Set("country", arg.Country).
		Set("entry_date", arg.EntryDate).
		Returning(selectStmt)

	row := r.Pool.QueryRow(ctx, b.String(), b.Args()...)
	var i Party
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.IsPublic,
		&i.MusicGenre,
		wkb.Scanner(&i.Location),
		&i.StreetAddress,
		&i.PostalCode,
		&i.State,
		&i.Country,
		&i.EntryDate,
		&i.MaxParticipants,
		&i.ParticipantsCount,
		&i.FavoriteCount,
	)
	return i, err
}

type UpdatePartyParams struct {
	ID            string
	Title         string
	Description   string
	MusicGenre    string
	Location      orb.Point
	StreetAddress string
	PostalCode    string
	State         string
	Country       string
	EntryDate     time.Time
}

func (r PartyRepository) UpdateParty(ctx context.Context, arg UpdatePartyParams) (Party, error) {
	sqlf.SetDialect(sqlf.PostgreSQL)
	b := sqlf.Update(TABLE_NAME)

	if arg.Title != "" {
		b = b.Set("title", arg.Title)
	}
	if arg.Description != "" {
		b = b.Set("description", arg.Description)
	}
	if arg.MusicGenre != "" {
		b = b.Set("music_genre", arg.MusicGenre)
	}
	if arg.Location.Lat() != 0 && arg.Location.Lon() != 0 {
		b = b.SetExpr("location", "ST_GeomFromEWKB(?)", ewkb.Value(arg.Location, 4326))
	}
	if arg.StreetAddress != "" {
		b = b.Set("street_address", arg.StreetAddress)
	}
	if arg.PostalCode != "" {
		b = b.Set("postal_code", arg.PostalCode)
	}
	if arg.State != "" {
		b = b.Set("state", arg.State)
	}
	if arg.Country != "" {
		b = b.Set("country", arg.Country)
	}
	startYear := arg.EntryDate.Year()
	if !(startYear == 1970) {
		b = b.Set("entry_date", arg.EntryDate)
	}

	b.
		Where("id = ?", arg.ID).
		Returning(selectStmt)

	row := r.Pool.QueryRow(ctx, b.String(), b.Args()...)
	var i Party
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.IsPublic,
		&i.MusicGenre,
		wkb.Scanner(&i.Location),
		&i.StreetAddress,
		&i.PostalCode,
		&i.State,
		&i.Country,
		&i.EntryDate,
		&i.MaxParticipants,
		&i.ParticipantsCount,
		&i.FavoriteCount,
	)

	return i, err
}

func (r PartyRepository) GetParty(ctx context.Context, id string) (Party, error) {
	sqlf.SetDialect(sqlf.PostgreSQL)
	b := sqlf.
		Select(selectStmt).
		From(TABLE_NAME).
		Where("id = ?", id)

	row := r.Pool.QueryRow(ctx, b.String(), b.Args()...)
	var i Party
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.IsPublic,
		&i.MusicGenre,
		wkb.Scanner(&i.Location),
		&i.StreetAddress,
		&i.PostalCode,
		&i.State,
		&i.Country,
		&i.EntryDate,
		&i.MaxParticipants,
		&i.ParticipantsCount,
		&i.FavoriteCount,
	)
	return i, err
}

type GetManyPartiesParams struct {
	Ids   []string
	Limit int32
}

func (r PartyRepository) GetManyParties(ctx context.Context, arg GetManyPartiesParams) ([]Party, error) {
	sqlf.SetDialect(sqlf.PostgreSQL)
	b := sqlf.
		Select(selectStmt).
		From(TABLE_NAME).
		Where("id = ANY(?)", arg.Ids).
		Limit(arg.Limit)

	rows, err := r.Pool.Query(ctx, b.String(), b.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Party
	for rows.Next() {
		var i Party
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.IsPublic,
			&i.MusicGenre,
			wkb.Scanner(&i.Location),
			&i.StreetAddress,
			&i.PostalCode,
			&i.State,
			&i.Country,
			&i.EntryDate,
			&i.MaxParticipants,
			&i.ParticipantsCount,
			&i.FavoriteCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

type GetPartiesByUserParams struct {
	UserID   string
	IsPublic bool
	Limit    int32
	Offset   int32
}

func (r PartyRepository) GetPartiesByUser(ctx context.Context, arg GetPartiesByUserParams) ([]Party, error) {
	sqlf.SetDialect(sqlf.PostgreSQL)
	b := sqlf.
		Select(selectStmt).
		From(TABLE_NAME).
		Where("user_id = ?", arg.UserID).
		Where("is_public = ?", arg.IsPublic).
		OrderBy("id desc")
	if arg.Limit == 0 {
		b = b.Limit(10)
	} else {
		b = b.Limit(arg.Limit)
	}
	b = b.Offset(arg.Offset)

	rows, err := r.Pool.Query(ctx, b.String(), b.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Party
	for rows.Next() {
		var i Party
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.IsPublic,
			&i.MusicGenre,
			wkb.Scanner(&i.Location),
			&i.StreetAddress,
			&i.PostalCode,
			&i.State,
			&i.Country,
			&i.EntryDate,
			&i.MaxParticipants,
			&i.ParticipantsCount,
			&i.FavoriteCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

type GeoSearchParams struct {
	Lat            float32
	Long           float32
	RadiusInDegree float32
	IsPublic       bool
	Limit          int32
	Offset         int32
}

func (r PartyRepository) GeoSearch(ctx context.Context, arg GeoSearchParams) ([]Party, error) {
	sqlf.SetDialect(sqlf.PostgreSQL)
	b := sqlf.
		Select(selectStmt).
		From(TABLE_NAME).
		Where("ST_DWithin(location, ST_SetSRID(ST_MakePoint(?, ?), 4326), ?)", arg.Long, arg.Lat, arg.RadiusInDegree).
		Where("is_public = ?", arg.IsPublic)

	if arg.Limit == 0 {
		b = b.Limit(10)
	} else {
		b = b.Limit(arg.Limit)
	}
	b = b.Offset(arg.Offset)

	rows, err := r.Pool.Query(ctx, b.String(), b.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Party
	for rows.Next() {
		var i Party
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Description,
			&i.IsPublic,
			&i.MusicGenre,
			wkb.Scanner(&i.Location),
			&i.StreetAddress,
			&i.PostalCode,
			&i.State,
			&i.Country,
			&i.EntryDate,
			&i.MaxParticipants,
			&i.ParticipantsCount,
			&i.FavoriteCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r PartyRepository) DeleteParty(ctx context.Context, arg DeletePartyParams) error {
	return r.querier.DeleteParty(ctx, arg)
}

func (r PartyRepository) IncreaseFavoriteCount(ctx context.Context, arg IncreaseFavoriteCountParams) error {
	return r.querier.IncreaseFavoriteCount(ctx, arg)
}

func (r PartyRepository) IncreaseParticipantsCount(ctx context.Context, arg IncreaseParticipantsCountParams) error {
	return r.querier.IncreaseParticipantsCount(ctx, arg)
}

func (r PartyRepository) DecreaseFavoriteCount(ctx context.Context, arg DecreaseFavoriteCountParams) error {
	return r.querier.DecreaseFavoriteCount(ctx, arg)
}

func (r PartyRepository) DecreaseParticipantsCount(ctx context.Context, arg DecreaseParticipantsCountParams) error {
	return r.querier.DecreaseParticipantsCount(ctx, arg)
}
