package service

import (
	"context"
	"sync"

	"github.com/clubo-app/clubben/participation-service/datastruct"
	"github.com/clubo-app/clubben/participation-service/repository"
	"github.com/clubo-app/clubben/protobuf/party"
)

type Invite interface {
	Invite(context.Context, repository.InviteParams) (datastruct.Invite, error)
	Decline(context.Context, repository.DeclineParams) error
	Accept(context.Context, repository.DeclineParams) (datastruct.Participant, error)
	GetUserInvites(context.Context, repository.GetUserInvitesParams) ([]datastruct.Invite, []byte, error)
}

type invite struct {
	r  repository.InviteRepository
	pc party.PartyServiceClient
	is Participant
}

func NewInviteService(r repository.InviteRepository, pc party.PartyServiceClient, is Participant) Invite {
	return invite{
		pc: pc,
		r:  r,
		is: is,
	}
}

func (i invite) Invite(ctx context.Context, params repository.InviteParams) (datastruct.Invite, error) {
	return i.r.Invite(ctx, params)
}

func (i invite) Decline(ctx context.Context, params repository.DeclineParams) error {
	return i.r.Decline(ctx, params)
}

func (i invite) Accept(ctx context.Context, params repository.DeclineParams) (datastruct.Participant, error) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	var err error
	var p datastruct.Participant

	go func() {
		defer wg.Done()

		tmp := i.r.Decline(ctx, params)
		if tmp != nil {
			err = tmp
		}
	}()

	go func() {
		defer wg.Done()

		var tmp error
		p, tmp = i.is.Join(ctx, JoinParams{
			UserId:    params.UserId,
			PartyId:   params.PartyId,
			InviterId: params.InviterId,
		})
		if tmp != nil {
			err = tmp
		}
	}()

	wg.Wait()

	return p, err
}

func (i invite) GetUserInvites(ctx context.Context, params repository.GetUserInvitesParams) ([]datastruct.Invite, []byte, error) {
	return i.r.GetUserInvites(ctx, params)
}
