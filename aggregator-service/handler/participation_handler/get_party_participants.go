package participationhandler

import (
	"strconv"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

func (h participationHandler) GetPartyParticipants(c *fiber.Ctx) error {
	pId := c.Params("pid")
	nextPage := c.Query("nextPage")
	limitStr := c.Query("limit")
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
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
			Limit:    int32(limit),
		})
	} else {
		pps, err = h.participationC.GetPartyParticipants(c.Context(), &participation.GetPartyParticipantsRequest{
			PartyId:  pId,
			NextPage: nextPage,
			Limit:    int32(limit),
		})
	}
	if err != nil {
		return utils.ToHTTPError(err)
	}

	userIds := make([]string, len(pps.Participants))
	for i, pp := range pps.Participants {
		userIds[i] = pp.UserId
	}
	profiles, _ := h.profileC.GetManyProfiles(c.Context(), &profile.GetManyProfilesRequest{Ids: utils.UniqueStringSlice(userIds)})

	aggP := make([]*datastruct.AggregatedProfile, len(profiles.Profiles))
	for i, profile := range profiles.Profiles {
		aggP[i] = datastruct.ProfileToAgg(profile)
	}

	res := datastruct.PagedAggregatedProfile{
		Profiles: aggP,
		NextPage: pps.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
