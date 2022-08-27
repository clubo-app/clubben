package participationhandler

import (
	"strconv"

	"github.com/clubo-app/aggregator-service/datastruct"
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/protobuf/participation"
	"github.com/clubo-app/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

func (h participationHandler) GetPartyParticipants(c *fiber.Ctx) error {
	pId := c.Params("pid")
	nextPage := c.Query("nextPage")
	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	if limit > 40 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 40")
	}

	requestedStr := c.Query("requested")
	requested, _ := strconv.ParseBool(requestedStr)

	var pps *participation.PagedPartyParticipants
	var err error
	if requested {
		pps, err = h.participationC.GetPartyRequests(c.Context(), &participation.GetPartyParticipantsRequest{
			PartyId:  pId,
			NextPage: nextPage,
			Limit:    uint32(limit),
		})
	} else {
		pps, err = h.participationC.GetPartyParticipants(c.Context(), &participation.GetPartyParticipantsRequest{
			PartyId:  pId,
			NextPage: nextPage,
			Limit:    uint32(limit),
		})
	}
	if err != nil {
		return utils.ToHTTPError(err)
	}

	userIds := make([]string, len(pps.Participants))
	for i, pp := range pps.Participants {
		userIds[i] = pp.UserId
	}
	profiles, _ := h.profileC.GetManyProfilesMap(c.Context(), &profile.GetManyProfilesRequest{Ids: utils.UniqueStringSlice(userIds)})

	aggP := make([]datastruct.AggregatedPartyParticipant, len(pps.Participants))
	for i, pp := range pps.Participants {
		p := datastruct.
			PartyParticipantToAgg(pp).
			AddParty(datastruct.AggregatedParty{Id: pp.PartyId})

		if profiles.Profiles[pp.UserId] != nil {
			p = p.AddUser(datastruct.ProfileToAgg(profiles.Profiles[pp.UserId]))
		}

		aggP[i] = p
	}

	res := datastruct.PagedAggregatedPartyParticipant{
		Participants: aggP,
		NextPage:     pps.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
