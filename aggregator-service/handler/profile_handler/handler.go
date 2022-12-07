package profilehandler

import (
	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

type profileHandler struct {
	profileClient  pg.ProfileServiceClient
	relationClient rg.RelationServiceClient
	authClient     pbauth.AuthServiceClient
}

type ProfileHandler interface {
	GetMe(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	UsernameTaken(c *fiber.Ctx) error
}

func NewUserHandler(profileClient pg.ProfileServiceClient, relationClient rg.RelationServiceClient, authClient pbauth.AuthServiceClient) ProfileHandler {
	return &profileHandler{
		profileClient:  profileClient,
		relationClient: relationClient,
		authClient:     authClient,
	}
}
