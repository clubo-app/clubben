package favoritehandler

import (
	"strconv"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h *favoriteHandler) GetFavoritePartiesByUser(c *fiber.Ctx) error {
	uId := c.Params("id")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	if limit > 20 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 20")
	}

	favParties, err := h.relationClient.GetFavoritePartiesByUser(c.Context(), &rg.GetFavoritePartiesByUserRequest{UserId: uId, NextPage: nextPage, Limit: limit})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	partyIds := make([]string, len(favParties.FavoriteParties))
	for i, fp := range favParties.FavoriteParties {
		partyIds[i] = fp.PartyId
	}

	if len(partyIds) == 0 {
		res := make([]string, 0)
		return c.Status(fiber.StatusOK).JSON(res)
	}

	parties, _ := h.partyClient.GetManyParties(c.Context(), &pbparty.GetManyPartiesRequest{Ids: partyIds})

	aggP := make([]*datastruct.AggregatedParty, len(favParties.FavoriteParties))
	for i, p := range parties.Parties {
		party := datastruct.PartyToAgg(p)
		party.IsFavorite = true

		aggP[i] = party
	}

	res := datastruct.PagedAggregatedParty{
		Parties:  aggP,
		NextPage: favParties.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
