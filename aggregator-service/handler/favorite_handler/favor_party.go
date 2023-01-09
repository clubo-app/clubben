package favoritehandler

import (
	"time"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h *favoriteHandler) FavorParty(c *fiber.Ctx) error {
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	pId := c.Params("pId")

	f, err := h.relationClient.FavorParty(c.Context(), &rg.PartyAndUserRequest{UserId: user.UserID, PartyId: pId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	res := datastruct.AggregatedFavoriteParty{
		UserId:      f.UserId,
		Party:       datastruct.PartyToAgg(&pbparty.Party{Id: f.PartyId}),
		FavoritedAt: f.FavoritedAt.AsTime().UTC().Format(time.RFC3339),
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
