package partyhandler

import (
	"github.com/clubo-app/clubben/protobuf/participation"
	pg "github.com/clubo-app/clubben/protobuf/party"
	prfg "github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	sg "github.com/clubo-app/clubben/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

type partyGatewayHandler struct {
	pc                  pg.PartyServiceClient
	prf                 prfg.ProfileServiceClient
	sc                  sg.StoryServiceClient
	rc                  rg.RelationServiceClient
	participationClient participation.ParticipationServiceClient
}

type PartyGatewayHandler interface {
	CreateParty(c *fiber.Ctx) error
	UpdateParty(c *fiber.Ctx) error
	DeleteParty(c *fiber.Ctx) error
	GetParty(c *fiber.Ctx) error
	GetPartyByUser(c *fiber.Ctx) error

	GetFavoritePartiesByUser(c *fiber.Ctx) error
	GetFavorisingUsersByParty(c *fiber.Ctx) error
	GeoSearch(c *fiber.Ctx) error
}

func NewPartyGatewayHandler(pc pg.PartyServiceClient, prf prfg.ProfileServiceClient, sc sg.StoryServiceClient, rc rg.RelationServiceClient, participationClient participation.ParticipationServiceClient) PartyGatewayHandler {
	return &partyGatewayHandler{
		pc:                  pc,
		prf:                 prf,
		sc:                  sc,
		rc:                  rc,
		participationClient: participationClient,
	}
}
