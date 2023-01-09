package favoritehandler

import (
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

type FavoriteHandler interface {
	FavorParty(c *fiber.Ctx) error
	DefavorParty(c *fiber.Ctx) error
	GetFavoritePartiesByUser(c *fiber.Ctx) error
	GetFavorisingUsersByParty(c *fiber.Ctx) error
}

type favoriteHandler struct {
	relationClient relation.RelationServiceClient
	partyClient    pbparty.PartyServiceClient
	profileClient  profile.ProfileServiceClient
}

func NewFavoriteHandler(
	relationClient relation.RelationServiceClient,
	partyClient pbparty.PartyServiceClient,
	profileClient profile.ProfileServiceClient,
) FavoriteHandler {
	return &favoriteHandler{
		relationClient: relationClient,
		partyClient:    partyClient,
		profileClient:  profileClient,
	}
}
