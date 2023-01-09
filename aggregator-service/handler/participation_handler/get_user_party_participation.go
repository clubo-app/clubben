package participationhandler

import (
	"strconv"
	"sync"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/participation"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h *participationHandler) GetUserPartyParticipation(c *fiber.Ctx) error {
	uId := c.Params("uid")
	nextPage := c.Query("nextPage")
	limitStr := c.Query("limit")
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	if limit > 40 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 40")
	}

	participation, err := h.participationClient.GetUserParticipations(c.Context(), &participation.GetUserParticipationsRequest{
		UserId:   uId,
		NextPage: nextPage,
		Limit:    int32(limit),
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	if len(participation.Participants) == 0 {
		return c.Status(fiber.StatusOK).JSON(datastruct.PagedAggregatedParty{Parties: make([]*datastruct.AggregatedParty, 0)})
	}

	partyIds := make([]string, len(participation.Participants))
	for i, p := range participation.Participants {
		partyIds[i] = p.PartyId
	}

	if len(partyIds) == 0 {
		res := make([]string, 0)
		return c.Status(fiber.StatusOK).JSON(res)
	}

	// we use two goroutines to parallize data fetching
	wg := new(sync.WaitGroup)
	wg.Add(2)

	var parties []*pbparty.Party
	go func() {
		defer wg.Done()
		partiesRes, _ := h.partyClient.GetManyParties(c.Context(), &pbparty.GetManyPartiesRequest{Ids: partyIds})
		parties = partiesRes.Parties
	}()

	var favoriteParties map[string]*relation.FavoriteParty
	go func() {
		defer wg.Done()
		fp, _ := h.relationClient.GetFavoritePartyManyParties(c.Context(), &relation.GetFavoritePartyManyPartiesRequest{
			UserId:   uId,
			PartyIds: partyIds,
		})
		favoriteParties = fp.FavoriteParties
	}()

	wg.Wait()

	aggParties := make([]*datastruct.AggregatedParty, len(parties))
	for i, p := range parties {
		party := datastruct.PartyToAgg(p)
		if _, ok := favoriteParties[party.Id]; ok {
			party.IsFavorite = true
		}

		aggParties[i] = party
	}

	res := datastruct.PagedAggregatedParty{
		Parties:  aggParties,
		NextPage: participation.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
