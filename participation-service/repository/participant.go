package repository

import (
	"context"
	"sync"
	"time"

	"github.com/clubo-app/clubben/participation-service/datastruct"
	"github.com/go-playground/validator"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx/table"
	"github.com/scylladb/gocqlx/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	PARTY_PARTICIPANTS         string = "party_participants"
	PARTY_PARTICIPANTS_BY_USER string = "party_participants_by_user"
	PARTY_PARTICIPATION_COUNT  string = "party_participation_count "
)

var participantMetadata = table.Metadata{
	Name:    PARTY_PARTICIPANTS,
	Columns: []string{"user_id", "party_id", "requested", "joined_at", "requested_at"},
	PartKey: []string{"party_id", "requested"},
	SortKey: []string{"user_id"},
}

var participationCount = table.Metadata{
	Name:    PARTY_PARTICIPATION_COUNT,
	Columns: []string{"party_id", "participation_count"},
	PartKey: []string{"party_id"},
}

type ParticipantRepository interface {
	Join(context.Context, UserPartyParams) (datastruct.Participant, error)
	Request(context.Context, UserPartyParams) (datastruct.Participant, error)
	Accept(context.Context, UserPartyParams) error
	Leave(context.Context, UserPartyParams) error
	GetPartyParticipants(context.Context, GetPartyParticipantsParams) ([]datastruct.Participant, []byte, error)
	GetPartyRequests(context.Context, GetPartyParticipantsParams) ([]datastruct.Participant, []byte, error)
	IncreaseParticipationCount(context.Context, string) error
	DecreaseParticipationCount(context.Context, string) error
	GetParticipationCount(context.Context, string) (int64, error)
}

type participantRepository struct {
	sess *gocqlx.Session
	val  *validator.Validate
}

type UserPartyParams struct {
	UserId  string
	PartyId string
}

func (r participantRepository) Join(ctx context.Context, params UserPartyParams) (p datastruct.Participant, err error) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		p = datastruct.Participant{
			UserId:    params.UserId,
			PartyId:   params.PartyId,
			JoinedAt:  time.Now(),
			Requested: false,
		}

		stmt, names := qb.
			Insert(PARTY_PARTICIPANTS).
			Unique().
			Columns(participantMetadata.Columns...).
			ToCql()

		err1 := r.sess.
			ContextQuery(ctx, stmt, names).
			BindStruct(p).
			ExecRelease()
		if err1 != nil {
			err = err1
		}
	}()

	go func() {
		defer wg.Done()
		err1 := r.IncreaseParticipationCount(ctx, params.PartyId)
		if err1 != nil {
			err = err1
		}
	}()
	wg.Wait()

	if err != nil {
		return datastruct.Participant{}, err
	}

	return p, nil
}

func (r participantRepository) Request(ctx context.Context, params UserPartyParams) (datastruct.Participant, error) {
	p := datastruct.Participant{
		UserId:      params.UserId,
		PartyId:     params.PartyId,
		Requested:   true,
		JoinedAt:    time.Now(),
		RequestedAt: time.Now(),
	}

	stmt, names := qb.
		Insert(PARTY_PARTICIPANTS).
		Unique().
		Columns(participantMetadata.Columns...).
		ToCql()

	err := r.sess.
		ContextQuery(ctx, stmt, names).
		BindStruct(p).
		ExecRelease()
	if err != nil {
		return datastruct.Participant{}, err
	}
	return p, nil
}

func (r participantRepository) Accept(ctx context.Context, params UserPartyParams) (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		stmt, names := qb.
			Update(PARTY_PARTICIPANTS).
			Where(qb.Eq("party_id")).
			Where(qb.EqNamed("requested", "old.requested")).
			Where(qb.Eq("user_id")).
			Set("requested").
			Set("joined_at").
			ToCql()

		err1 := r.sess.
			ContextQuery(ctx, stmt, names).
			BindMap((qb.M{
				"party_id":      params.PartyId,
				"user_id":       params.UserId,
				"old.requested": true,
				"requested":     false,
				"joined_at":     time.Now(),
			})).
			ExecRelease()
		if err1 != nil {
			err = err1
		}
	}()
	go func() {
		defer wg.Done()
		err1 := r.IncreaseParticipationCount(ctx, params.PartyId)
		if err1 != nil {
			err = err1
		}
	}()
	wg.Wait()

	return err
}

func (r participantRepository) Leave(ctx context.Context, params UserPartyParams) (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()

		stmt, names := qb.
			Delete(PARTY_PARTICIPANTS).
			Where(qb.Eq("user_id")).
			Where(qb.Eq("party_id")).
			ToCql()

		err1 := r.sess.
			ContextQuery(ctx, stmt, names).
			BindMap((qb.M{
				"user_id":  params.UserId,
				"party_id": params.PartyId,
			})).
			ExecRelease()
		if err1 != nil {
			err = err1
		}
	}()
	go func() {
		defer wg.Done()
		err1 := r.DecreaseParticipationCount(ctx, params.PartyId)
		if err1 != nil {
			err = err1
		}
	}()

	wg.Wait()
	if err != nil {
		return err
	}
	return nil
}

type GetPartyParticipantsParams struct {
	PId   string
	Page  []byte
	Limit int
}

func (r participantRepository) GetPartyParticipants(ctx context.Context, params GetPartyParticipantsParams) (res []datastruct.Participant, nextPage []byte, err error) {
	stmt, names := qb.
		Select(PARTY_PARTICIPANTS).
		Where(qb.Eq("party_id")).
		Where(qb.Eq("requested")).
		ToCql()

	q := r.sess.
		ContextQuery(ctx, stmt, names).
		BindMap((qb.M{
			"party_id":  params.PId,
			"requested": false,
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
		return []datastruct.Participant{}, nil, status.Error(codes.Internal, "No Participants found")
	}

	return res, iter.PageState(), nil
}

func (r participantRepository) GetPartyRequests(ctx context.Context, params GetPartyParticipantsParams) (res []datastruct.Participant, nextPage []byte, err error) {
	stmt, names := qb.
		Select(PARTY_PARTICIPANTS).
		Where(qb.Eq("party_id")).
		Where(qb.Eq("requested")).
		ToCql()

	q := r.sess.
		ContextQuery(ctx, stmt, names).
		BindMap((qb.M{
			"party_id":  params.PId,
			"requested": true,
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
		return []datastruct.Participant{}, nil, status.Error(codes.Internal, "No Requests found")
	}

	return res, iter.PageState(), nil
}

func (r participantRepository) IncreaseParticipationCount(ctx context.Context, pId string) error {
	countStmt, countNames := qb.
		Update(PARTY_PARTICIPATION_COUNT).
		Where(qb.Eq("party_id")).
		Add("participation_count").
		ToCql()

	err := r.sess.
		ContextQuery(ctx, countStmt, countNames).
		BindMap((qb.M{
			"participation_count": 1,
			"party_id":            pId,
		})).
		ExecRelease()
	if err != nil {
		return err
	}
	return nil
}

func (r participantRepository) DecreaseParticipationCount(ctx context.Context, pId string) error {
	countStmt, countNames := qb.
		Update(PARTY_PARTICIPATION_COUNT).
		Where(qb.Eq("party_id")).
		Remove("participation_count").
		ToCql()

	err := r.sess.
		ContextQuery(ctx, countStmt, countNames).
		BindMap((qb.M{
			"participation_count": 1,
			"party_id":            pId,
		})).
		ExecRelease()
	if err != nil {
		return err
	}
	return nil
}

func (r participantRepository) GetParticipationCount(ctx context.Context, pId string) (int64, error) {
	var res datastruct.ParticipationCount

	stmt, names := qb.
		Select(PARTY_PARTICIPATION_COUNT).
		Columns(participationCount.Columns...).
		Where(qb.Eq("party_id")).
		ToCql()

	err := r.sess.
		ContextQuery(ctx, stmt, names).
		BindMap((qb.M{
			"party_id": pId,
		})).
		GetRelease(&res)
	if err != nil {
		return 0, err
	}

	return res.ParticipationCount, err
}
