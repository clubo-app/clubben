package favoritehandler

import (
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h favoriteHandler) DefavorParty(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	pId := c.Params("pId")

	ok, err := h.relationClient.DefavorParty(c.Context(), &rg.FavorPartyRequest{UserId: user.Sub, PartyId: pId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(ok)
}
