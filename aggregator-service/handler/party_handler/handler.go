package partyhandler

import (
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	"github.com/clubo-app/clubben/protobuf/participation"
	prfg "github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	sg "github.com/clubo-app/clubben/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

type partyHandler struct {
	partyClient         pbparty.PartyServiceClient
	profileClient       prfg.ProfileServiceClient
	storyClient         sg.StoryServiceClient
	relationClient      rg.RelationServiceClient
	participationClient participation.ParticipationServiceClient
}

type PartyHandler interface {
	CreateParty(c *fiber.Ctx) error
	UpdateParty(c *fiber.Ctx) error
	DeleteParty(c *fiber.Ctx) error
	GetParty(c *fiber.Ctx) error
	GetPartyByUser(c *fiber.Ctx) error

	GeoSearch(c *fiber.Ctx) error
}

func NewPartyHandler(
	partyClient pbparty.PartyServiceClient,
	profileClient prfg.ProfileServiceClient,
	storyClient sg.StoryServiceClient,
	relationClient rg.RelationServiceClient,
	participationClient participation.ParticipationServiceClient,
) PartyHandler {
	return &partyHandler{
		partyClient:         partyClient,
		profileClient:       profileClient,
		storyClient:         storyClient,
		relationClient:      relationClient,
		participationClient: participationClient,
	}
}
