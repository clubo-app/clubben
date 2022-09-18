package participationhandler

import (
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

type participationHandler struct {
	profileClient       profile.ProfileServiceClient
	partyClient         party.PartyServiceClient
	participationClient participation.ParticipationServiceClient
	relationClient      relation.RelationServiceClient
}

type ParticipationHandler interface {
	InviteToParty(c *fiber.Ctx) error
	DeclinePartyInvite(c *fiber.Ctx) error
	AcceptPartyInvite(c *fiber.Ctx) error
	GetUserInvites(c *fiber.Ctx) error
	JoinParty(c *fiber.Ctx) error
	LeaveParty(c *fiber.Ctx) error
	GetPartyParticipants(c *fiber.Ctx) error

	GetUserPartyParticipation(c *fiber.Ctx) error
	GetFriendsPartyParticipation(c *fiber.Ctx) error
}

func NewParticipationHandler(
	participationClient participation.ParticipationServiceClient,
	partyClient party.PartyServiceClient,
	profileClient profile.ProfileServiceClient,
	relationClient relation.RelationServiceClient,
) ParticipationHandler {
	return &participationHandler{
		participationClient: participationClient,
		partyClient:         partyClient,
		profileClient:       profileClient,
		relationClient:      relationClient,
	}
}
