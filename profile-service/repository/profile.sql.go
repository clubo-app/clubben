// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: profile.sql

package repository

import (
	"context"
	"database/sql"
)

const createProfile = `-- name: CreateProfile :one
INSERT INTO profiles (
    id,
    username,
    firstname,
    lastname,
    avatar
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, username, firstname, lastname, avatar
`

type CreateProfileParams struct {
	ID        string
	Username  string
	Firstname string
	Lastname  sql.NullString
	Avatar    sql.NullString
}

func (q *Queries) CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error) {
	row := q.db.QueryRow(ctx, createProfile,
		arg.ID,
		arg.Username,
		arg.Firstname,
		arg.Lastname,
		arg.Avatar,
	)
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

const deleteProfile = `-- name: DeleteProfile :exec
DELETE FROM profiles
WHERE id = $1
`

func (q *Queries) DeleteProfile(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteProfile, id)
	return err
}

const getManyProfiles = `-- name: GetManyProfiles :many
SELECT id, username, firstname, lastname, avatar FROM profiles
WHERE id=ANY($1::text[])
LIMIT $2
`

type GetManyProfilesParams struct {
	Ids   []string
	Limit int32
}

func (q *Queries) GetManyProfiles(ctx context.Context, arg GetManyProfilesParams) ([]Profile, error) {
	rows, err := q.db.Query(ctx, getManyProfiles, arg.Ids, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Profile
	for rows.Next() {
		var i Profile
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Firstname,
			&i.Lastname,
			&i.Avatar,
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

const getProfile = `-- name: GetProfile :one
SELECT id, username, firstname, lastname, avatar FROM profiles
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProfile(ctx context.Context, id string) (Profile, error) {
	row := q.db.QueryRow(ctx, getProfile, id)
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

const getProfileByUsername = `-- name: GetProfileByUsername :one
SELECT id, username, firstname, lastname, avatar FROM profiles
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetProfileByUsername(ctx context.Context, username string) (Profile, error) {
	row := q.db.QueryRow(ctx, getProfileByUsername, username)
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

const usernameTaken = `-- name: UsernameTaken :one
select exists(select 1 from profiles where username=$1) AS "exists"
`

func (q *Queries) UsernameTaken(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRow(ctx, usernameTaken, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
