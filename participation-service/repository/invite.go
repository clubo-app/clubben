package repository

import (
	"context"
	"log"
	"time"

	"github.com/clubo-app/clubben/participation-service/datastruct"
	"github.com/go-playground/validator"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx/table"
	"github.com/scylladb/gocqlx/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const PARTY_INVITES string = "party_invites"

var inviteMetadata = table.Metadata{
	Name:    PARTY_INVITES,
	Columns: []string{"user_id", "party_id", "inviter_id", "valid_until"},
	PartKey: []string{"user_id", "inviter_id", "party_id"},
}

type InviteRepository interface {
	Invite(context.Context, InviteParams) (datastruct.Invite, error)
	Decline(context.Context, DeclineParams) error
	GetUserInvites(context.Context, GetUserInvitesParams) ([]datastruct.Invite, []byte, error)
}

type inviteRepository struct {
	sess *gocqlx.Session
	val  *validator.Validate
}

type InviteParams struct {
	UserId    string
	InviterId string
	PartyId   string
	ValidFor  time.Duration
}

func (r inviteRepository) Invite(ctx context.Context, params InviteParams) (datastruct.Invite, error) {
	i := datastruct.Invite{
		UserId:     params.UserId,
		InviterId:  params.InviterId,
		PartyId:    params.PartyId,
		ValidUntil: time.Now().Add(params.ValidFor),
	}
	err := r.val.StructCtx(ctx, i)
	if err != nil {
		return datastruct.Invite{}, err
	}

	stmt, names := qb.
		Insert(PARTY_INVITES).
		Unique().
		Columns(inviteMetadata.Columns...).
		TTL(params.ValidFor).
		ToCql()

	err = r.sess.
		ContextQuery(ctx, stmt, names).
		BindStruct(i).
		ExecRelease()
	if err != nil {
		return datastruct.Invite{}, err
	}

	return i, nil
}

type DeclineParams struct {
	UserId    string
	InviterId string
	PartyId   string
}

func (r inviteRepository) Decline(ctx context.Context, params DeclineParams) error {
	stmt, names := qb.
		Delete(PARTY_INVITES).
		Where(qb.Eq("user_id")).
		Where(qb.Eq("inviter_id")).
		Where(qb.Eq("party_id")).
		ToCql()

	log.Printf("%+v\n", params)

	err := r.sess.ContextQuery(ctx, stmt, names).
		BindMap((qb.M{
			"user_id":    params.UserId,
			"inviter_id": params.InviterId,
			"party_id":   params.PartyId,
		})).
		ExecRelease()
	if err != nil {
		return err
	}

	return nil
}

type GetUserInvitesParams struct {
	UId   string
	Page  []byte
	Limit int
}

func (r inviteRepository) GetUserInvites(ctx context.Context, params GetUserInvitesParams) (res []datastruct.Invite, nextPage []byte, err error) {
	stmt, names := qb.
		Select(PARTY_INVITES).
		Where(qb.Eq("user_id")).
		ToCql()

	q := r.sess.
		ContextQuery(ctx, stmt, names).
		BindMap((qb.M{
			"user_id": params.UId,
		}))

	q.PageState(params.Page)
	if params.Limit == 0 {
		q.PageSize(20)
	} else {
		q.PageSize(params.Limit)
	}

	iter := q.Iter()
	err = iter.Select(&res)
	if err != nil {
		return []datastruct.Invite{}, nil, status.Error(codes.Internal, "No invites found")
	}

	return res, iter.PageState(), nil
}
