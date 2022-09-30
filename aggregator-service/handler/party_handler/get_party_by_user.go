package partyhandler

import (
	"strconv"
	"sync"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h partyHandler) GetPartyByUser(c *fiber.Ctx) error {
	uId := c.Params("id")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	if limit > 20 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 20")
	}

	isPublicStr := c.Query("is_public")
	isPublic, _ := strconv.ParseBool(isPublicStr)

	offsetStr := c.Query("offset")
	offset, _ := strconv.ParseInt(offsetStr, 10, 32)

	parties, err := h.partyClient.GetByUser(c.Context(), &party.GetByUserRequest{
		UserId:   uId,
		Offset:   int32(offset),
		Limit:    int32(limit),
		IsPublic: isPublic,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	// we use two goroutines to parallize data fetching
	wg := new(sync.WaitGroup)
	wg.Add(2)

	var profileRes *profile.Profile
	go func() {
		defer wg.Wait()
		profileRes, _ = h.profileClient.GetProfile(c.Context(), &profile.GetProfileRequest{Id: uId})
	}()

	var favoriteParties map[string]*relation.FavoriteParty
	go func() {
		defer wg.Wait()
		partyIds := make([]string, len(parties.Parties))
		for i, p := range parties.Parties {
			partyIds[i] = p.Id
		}

		fp, _ := h.relationClient.GetFavoritePartyManyParties(c.Context(), &relation.GetFavoritePartyManyPartiesRequest{
			UserId:   uId,
			PartyIds: partyIds,
		})
		favoriteParties = fp.FavoriteParties
	}()

	wg.Wait()

	aggP := make([]*datastruct.AggregatedParty, len(parties.Parties))
	for i, p := range parties.Parties {
		party := datastruct.
			PartyToAgg(p).
			AddCreator(datastruct.ProfileToAgg(profileRes))
		if _, ok := favoriteParties[p.Id]; ok {
			party.IsFavorite = true
		}

		aggP[i] = party
	}

	return c.Status(fiber.StatusOK).JSON(aggP)
}
