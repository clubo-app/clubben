package relationhandler

import (
	"time"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	"github.com/clubo-app/clubben/protobuf/party"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h relationGatewayHandler) FavorParty(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	pId := c.Params("id")

	f, err := h.rc.FavorParty(c.Context(), &rg.FavorPartyRequest{UserId: user.Sub, PartyId: pId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	res := datastruct.AggregatedFavoriteParty{
		UserId:      f.UserId,
		Party:       datastruct.PartyToAgg(&party.Party{Id: f.PartyId}),
		FavoritedAt: f.FavoritedAt.AsTime().UTC().Format(time.RFC3339),
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
