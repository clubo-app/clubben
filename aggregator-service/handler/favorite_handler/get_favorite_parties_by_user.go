package favoritehandler

import (
	"strconv"
	"time"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/party"
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

	parties, _ := h.partyClient.GetManyPartiesMap(c.Context(), &party.GetManyPartiesRequest{Ids: partyIds})
	if parties == nil {
		res := datastruct.PagedAggregatedFavoriteParty{
			FavoriteParties: []datastruct.AggregatedFavoriteParty{},
			NextPage:        favParties.NextPage,
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}

	aggFP := make([]datastruct.AggregatedFavoriteParty, len(favParties.FavoriteParties))
	for i, fp := range favParties.FavoriteParties {
		aggFP[i] = datastruct.AggregatedFavoriteParty{
			UserId:      fp.UserId,
			Party:       datastruct.PartyToAgg(parties.Parties[fp.PartyId]),
			FavoritedAt: fp.FavoritedAt.AsTime().UTC().Format(time.RFC3339),
		}
	}

	res := datastruct.PagedAggregatedFavoriteParty{
		FavoriteParties: aggFP,
		NextPage:        favParties.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
