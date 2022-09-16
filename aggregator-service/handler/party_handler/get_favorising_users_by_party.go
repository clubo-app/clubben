package partyhandler

import (
	"strconv"

	"github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
)

func (h partyGatewayHandler) GetFavorisingUsersByParty(c *fiber.Ctx) error {
	pId := c.Params("pId")
	nextPage := c.Query("nextPage")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	if limit > 20 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 20")
	}

	fpRes, err := h.rc.GetFavorisingUsersByParty(c.Context(), &rg.GetFavorisingUsersByPartyRequest{PartyId: pId, NextPage: nextPage, Limit: limit})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	ids := make([]string, len(fpRes.FavoriteParties))
	for i, fp := range fpRes.FavoriteParties {
		ids[i] = fp.UserId
	}

	profiles, _ := h.prf.GetManyProfiles(c.Context(), &profile.GetManyProfilesRequest{Ids: ids})
	if profiles == nil {
		res := datastruct.PagedAggregatedProfile{
			Profiles: []*datastruct.AggregatedProfile{},
			NextPage: fpRes.NextPage,
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}

	aggP := make([]*datastruct.AggregatedProfile, len(profiles.Profiles))
	for i, profile := range profiles.Profiles {
		aggP[i] = datastruct.ProfileToAgg(profile)
	}

	res := datastruct.PagedAggregatedProfile{
		Profiles: aggP,
		NextPage: fpRes.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
