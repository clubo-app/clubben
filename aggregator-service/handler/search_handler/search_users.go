package searchhandler

import (
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/search"
	"github.com/gofiber/fiber/v2"
)

func (h searchGatewayHandler) SearchUsers(c *fiber.Ctx) error {
	q := c.Params("query")
	res, err := h.sc.SearchUsers(c.Context(), &search.SearchUsersRequest{Query: q})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
