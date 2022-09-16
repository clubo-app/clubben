package service

import (
	"context"
	"log"

	"github.com/clubo-app/clubben/participation-service/datastruct"
	"github.com/clubo-app/clubben/participation-service/repository"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/clubo-app/packages/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Participant interface {
	Join(context.Context, JoinParams) (datastruct.Participant, error)
	Request(context.Context, repository.UserPartyParams) (datastruct.Participant, error)
	Accept(context.Context, AcceptRequestParams) error
	Leave(context.Context, repository.UserPartyParams) error
	GetPartyParticipant(context.Context, repository.UserPartyParams) (datastruct.Participant, error)
	GetPartyParticipants(context.Context, repository.GetPartyParticipantsParams) ([]datastruct.Participant, []byte, error)
	GetPartyRequests(context.Context, repository.GetPartyParticipantsParams) ([]datastruct.Participant, []byte, error)
}

type participant struct {
	r  repository.ParticipantRepository
	pc party.PartyServiceClient
}

func NewParticipantService(r repository.ParticipantRepository, pc party.PartyServiceClient) Participant {
	return &participant{
		r:  r,
		pc: pc,
	}
}

type JoinParams struct {
	UserId    string
	PartyId   string
	InviterId string
}

func (p participant) Join(ctx context.Context, params JoinParams) (datastruct.Participant, error) {
	party, err := p.pc.GetParty(ctx, &party.GetPartyRequest{
		PartyId: params.PartyId,
	})
	if err != nil || party == nil {
		log.Println("Join Error: ", err)
		return datastruct.Participant{}, status.Error(codes.InvalidArgument, "Party not found")
	}
	if params.UserId == party.UserId {
		return datastruct.Participant{}, status.Error(codes.InvalidArgument, "You can't join your own Party")
	}

	log.Println("Join Got Party: ", party)

	count, _ := p.r.GetParticipationCount(ctx, params.PartyId)
	if err != nil {
		log.Println("ParticipationCount: ", count)
		log.Println("Join Error ParticipationCount: ", err)
		return datastruct.Participant{}, utils.HandleError(err)
	}

	// MaxParticipants of 0 mean there is no restriction for Participants
	if party.MaxParticipants != 0 && count >= int64(party.MaxParticipants) {
		return datastruct.Participant{}, status.Error(codes.PermissionDenied, "Party is already full")
	}

	// we request access to a party when it is private and when the user was not invited by the creator of the Party
	if !party.IsPublic && params.InviterId != party.UserId {
		return p.r.Request(ctx, repository.UserPartyParams{UserId: params.UserId, PartyId: params.PartyId})
	}

	return p.r.Join(ctx, repository.UserPartyParams{UserId: params.UserId, PartyId: params.PartyId})
}

func (p participant) Request(ctx context.Context, params repository.UserPartyParams) (datastruct.Participant, error) {
	return p.r.Request(ctx, params)
}

type AcceptRequestParams struct {
	PartyId    string
	UserId     string
	AccepterId string
}

func (p participant) Accept(ctx context.Context, params AcceptRequestParams) error {
	party, err := p.pc.GetParty(ctx, &party.GetPartyRequest{
		PartyId: params.PartyId,
	})
	if err != nil || party == nil {
		log.Println("Accept Error: ", err)
		return status.Error(codes.InvalidArgument, "Party not found")
	}
	if params.AccepterId != party.UserId {
		return status.Error(codes.PermissionDenied, "You are not the creator of the Party")
	}

	count, err := p.r.GetParticipationCount(ctx, params.PartyId)
	if count >= int64(party.MaxParticipants) {
		return status.Error(codes.PermissionDenied, "Party is already full")
	}

	return p.r.Accept(ctx, repository.UserPartyParams{
		UserId:  params.UserId,
		PartyId: params.PartyId,
	})
}

func (p participant) Leave(ctx context.Context, params repository.UserPartyParams) error {
	return p.r.Leave(ctx, params)
}

func (p participant) GetPartyParticipant(ctx context.Context, params repository.UserPartyParams) (datastruct.Participant, error) {
	return p.r.GetPartyParticipant(ctx, params)
}

func (p participant) GetPartyParticipants(ctx context.Context, params repository.GetPartyParticipantsParams) ([]datastruct.Participant, []byte, error) {
	return p.r.GetPartyParticipants(ctx, params)
}

func (p participant) GetPartyRequests(ctx context.Context, params repository.GetPartyParticipantsParams) ([]datastruct.Participant, []byte, error) {
	return p.r.GetPartyRequests(ctx, params)
}
