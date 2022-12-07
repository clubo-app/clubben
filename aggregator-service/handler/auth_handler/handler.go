package authhandler

import (
	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authClient    pbauth.AuthServiceClient
	profileClient pg.ProfileServiceClient
}

type AuthHandler interface {
	Register(c *fiber.Ctx) error
}

func NewAuthHandler(authClient pbauth.AuthServiceClient, profileClient pg.ProfileServiceClient) AuthHandler {
	return &authHandler{
		authClient:    authClient,
		profileClient: profileClient,
	}
}
