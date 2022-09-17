package relationhandler

import (
	pg "github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

type relationHandler struct {
	profileClient  pg.ProfileServiceClient
	relationClient rg.RelationServiceClient
}

type RelationHandler interface {
	CreateFriendRequest(c *fiber.Ctx) error
	AcceptFriendRequest(c *fiber.Ctx) error
	DeclineFriendRequest(c *fiber.Ctx) error
	RemoveFriend(c *fiber.Ctx) error
	GetFriends(c *fiber.Ctx) error
}

func NewRelationHandler(relationClient rg.RelationServiceClient, profileClient pg.ProfileServiceClient) RelationHandler {
	return &relationHandler{
		relationClient: relationClient,
		profileClient:  profileClient,
	}
}
