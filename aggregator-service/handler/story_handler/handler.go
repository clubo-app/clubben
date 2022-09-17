package storyhandler

import (
	pg "github.com/clubo-app/clubben/protobuf/party"
	prfg "github.com/clubo-app/clubben/protobuf/profile"
	sg "github.com/clubo-app/clubben/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

type storyHandler struct {
	storyClient   sg.StoryServiceClient
	partyClient   pg.PartyServiceClient
	profileClient prfg.ProfileServiceClient
}

type StoryHandler interface {
	CreateStory(c *fiber.Ctx) error
	GetStoryByParty(c *fiber.Ctx) error
	GetStoryByUser(c *fiber.Ctx) error
	GetPastStoryByUser(c *fiber.Ctx) error
	DeleteStory(c *fiber.Ctx) error
	PresignURL(c *fiber.Ctx) error
}

func NewStoryHandler(storyClient sg.StoryServiceClient, profileClient prfg.ProfileServiceClient, partyClient pg.PartyServiceClient) StoryHandler {
	return &storyHandler{
		storyClient:   storyClient,
		profileClient: profileClient,
		partyClient:   partyClient,
	}
}
