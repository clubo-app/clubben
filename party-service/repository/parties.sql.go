// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: parties.sql

package repository

import (
	"context"
)

const deleteParty = `-- name: DeleteParty :exec
DELETE FROM parties
WHERE id = $1 AND user_id = $2
RETURNING id, user_id, title, is_public, max_participants, location, street_address, postal_code, state, country, start_date, end_date
`

type DeletePartyParams struct {
	ID     string
	UserID string
}

func (q *Queries) DeleteParty(ctx context.Context, arg DeletePartyParams) error {
	_, err := q.db.Exec(ctx, deleteParty, arg.ID, arg.UserID)
	return err
}
