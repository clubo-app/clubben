package searchhandler

import (
	"github.com/clubo-app/clubben/protobuf/search"
	"github.com/gofiber/fiber/v2"
)

type searchGatewayHandler struct {
	sc search.SearchServiceClient
}

type SearchGatewayHandler interface {
	SearchUsers(c *fiber.Ctx) error
}

func NewSearchGatewayHandler(sc search.SearchServiceClient) SearchGatewayHandler {
	return &searchGatewayHandler{
		sc: sc,
	}
}
