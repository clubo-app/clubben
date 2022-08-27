package repository

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/leporo/sqlf"
)

const (
	TABLE_NAME = "profiles"
)

type ProfileRepository struct {
	pool    *pgxpool.Pool
	querier Querier
}

func NewProfileRepository(dbUser, dbPW, dbName, dbHost string, dbPort uint16) (*ProfileRepository, error) {
	urlStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPW, dbHost, fmt.Sprint(dbPort), dbName)
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

	return &ProfileRepository{
		pool:    pool,
		querier: New(pool),
	}, nil
}

func (d ProfileRepository) Close() {
	d.pool.Close()
}

const columns = "id, username, firstname, lastname, avatar"

func (r ProfileRepository) CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error) {
	return r.querier.CreateProfile(ctx, arg)
}

type UpdateProfileParams struct {
	ID        string
	Username  string
	Firstname string
	Lastname  string
	Avatar    string
}

func (r ProfileRepository) UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error) {
	sqlf.SetDialect(sqlf.PostgreSQL)
	b := sqlf.Update(TABLE_NAME)

	if arg.Username != "" {
		b = b.Set("username", arg.Username)
	}
	if arg.Firstname != "" {
		b = b.Set("firstname", arg.Firstname)
	}
	if arg.Lastname != "" {
		b = b.Set("lastname", arg.Lastname)
	}
	if arg.Avatar != "" {
		b = b.Set("avatar", arg.Avatar)
	}

	b.
		Where("id = ?", arg.ID).
		Returning(columns)

	row := r.pool.QueryRow(ctx, b.String(), b.Args()...)
	var i Profile
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Firstname,
		&i.Lastname,
		&i.Avatar,
	)
	return i, err
}

func (r ProfileRepository) DeleteProfile(ctx context.Context, id string) error {
	return r.querier.DeleteProfile(ctx, id)
}

func (r ProfileRepository) GetProfile(ctx context.Context, id string) (Profile, error) {
	return r.querier.GetProfile(ctx, id)
}

func (r ProfileRepository) GetProfileByUsername(ctx context.Context, username string) (Profile, error) {
	return r.querier.GetProfileByUsername(ctx, username)
}

func (r ProfileRepository) GetManyProfiles(ctx context.Context, arg GetManyProfilesParams) ([]Profile, error) {
	return r.querier.GetManyProfiles(ctx, arg)
}

func (r ProfileRepository) UsernameTaken(ctx context.Context, username string) (bool, error) {
	return r.querier.UsernameTaken(ctx, username)
}
