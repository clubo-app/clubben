package profilehandler

import (
	ag "github.com/clubo-app/clubben/protobuf/auth"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

type profileHandler struct {
	profileClient  pg.ProfileServiceClient
	relationClient rg.RelationServiceClient
	authClient     ag.AuthServiceClient
}

type ProfileHandler interface {
	GetMe(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	UsernameTaken(c *fiber.Ctx) error
}

func NewUserHandler(profileClient pg.ProfileServiceClient, relationClient rg.RelationServiceClient, authClient ag.AuthServiceClient) ProfileHandler {
	return &profileHandler{
		profileClient:  profileClient,
		relationClient: relationClient,
		authClient:     authClient,
	}
}
