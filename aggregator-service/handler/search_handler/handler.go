package searchhandler

import (
	"github.com/clubo-app/clubben/protobuf/search"
	"github.com/gofiber/fiber/v2"
)

type searchHandler struct {
	searchClient search.SearchServiceClient
}

type SearchHandler interface {
	SearchUsers(c *fiber.Ctx) error
	SearchParties(c *fiber.Ctx) error
}

func NewSearchHandler(searchClient search.SearchServiceClient) SearchHandler {
	return &searchHandler{
		searchClient: searchClient,
	}
}
