package participationhandler

import (
	"github.com/clubo-app/protobuf/participation"
	"github.com/clubo-app/protobuf/party"
	"github.com/clubo-app/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

type participationHandler struct {
	profileC       profile.ProfileServiceClient
	partyC         party.PartyServiceClient
	participationC participation.ParticipationServiceClient
}

type ParticipationHandler interface {
	InviteToParty(c *fiber.Ctx) error
	DeclinePartyInvite(c *fiber.Ctx) error
	AcceptPartyInvite(c *fiber.Ctx) error
	GetUserInvites(c *fiber.Ctx) error
	JoinParty(c *fiber.Ctx) error
	LeaveParty(c *fiber.Ctx) error
	GetPartyParticipants(c *fiber.Ctx) error
}

func NewParticipationHandler(pc participation.ParticipationServiceClient, partyC party.PartyServiceClient, profileC profile.ProfileServiceClient) ParticipationHandler {
	return &participationHandler{
		participationC: pc,
		partyC:         partyC,
		profileC:       profileC,
	}
}
