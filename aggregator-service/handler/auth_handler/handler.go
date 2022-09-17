package authhandler

import (
	ag "github.com/clubo-app/clubben/protobuf/auth"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authClient    ag.AuthServiceClient
	profileClient pg.ProfileServiceClient
}

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	GoogleLogin(c *fiber.Ctx) error
	VerifyEmail(c *fiber.Ctx) error
	RefreshAccessToken(c *fiber.Ctx) error
}

func NewAuthHandler(authClient ag.AuthServiceClient, profileClient pg.ProfileServiceClient) AuthHandler {
	return &authHandler{
		authClient:    authClient,
		profileClient: profileClient,
	}
}
