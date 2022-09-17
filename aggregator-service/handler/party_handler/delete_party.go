package partyhandler

import (
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/gofiber/fiber/v2"
)

func (h partyHandler) DeleteParty(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	pId := c.Params("id")

	res, err := h.partyClient.DeleteParty(c.Context(), &party.DeletePartyRequest{RequesterId: user.Sub, PartyId: pId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
