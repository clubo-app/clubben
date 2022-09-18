package participationhandler

import (
	"strconv"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/gofiber/fiber/v2"
)

func (h participationHandler) GetUserPartyParticipation(c *fiber.Ctx) error {
	uId := c.Params("uid")
	nextPage := c.Query("nextPage")
	limitStr := c.Query("limit")
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	if limit > 40 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 40")
	}

	participation, err := h.participationClient.GetUserParticipation(c.Context(), &participation.GetUserParticipationRequest{
		UserId:   uId,
		NextPage: nextPage,
		Limit:    int32(limit),
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	if len(participation.Participants) == 0 {
		return c.Status(fiber.StatusOK).JSON(make([]string, 0))
	}

	partyIds := make([]string, len(participation.Participants))
	for i, p := range participation.Participants {
		partyIds[i] = p.PartyId
	}

	parties, err := h.partyClient.GetManyParties(c.Context(), &party.GetManyPartiesRequest{
		Ids: partyIds,
	})

	res := make([]*datastruct.AggregatedParty, len(parties.Parties))
	for i, party := range parties.Parties {
		res[i] = datastruct.PartyToAgg(party).AddCreator(&datastruct.AggregatedProfile{Id: uId})
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
