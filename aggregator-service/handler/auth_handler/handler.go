package authhandler

import (
	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	pbrelation "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authClient     pbauth.AuthServiceClient
	profileClient  pg.ProfileServiceClient
	relationClient pbrelation.RelationServiceClient
}

type AuthHandler interface {
	GetMe(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	RegisterAnonymously(c *fiber.Ctx) error
  GetToken(c *fiber.Ctx) error
}

func NewAuthHandler(
	authClient pbauth.AuthServiceClient,
	profileClient pg.ProfileServiceClient,
	relationClient pbrelation.RelationServiceClient,
) AuthHandler {
	return &authHandler{
		authClient:     authClient,
		profileClient:  profileClient,
		relationClient: relationClient,
	}
}
