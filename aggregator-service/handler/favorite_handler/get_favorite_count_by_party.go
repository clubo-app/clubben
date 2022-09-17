package favoritehandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h favoriteHandler) GetFavoriteCountByParty(c *fiber.Ctx) error {
	pId := c.Params("pId")

	count, err := h.relationClient.GetFavoritePartyCount(c.Context(), &relation.GetFavoritePartyCountRequest{
		PartyId: pId,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(datastruct.PartyFavoriteCount{FavoriteCount: int(count.FavoriteCount)})
}
