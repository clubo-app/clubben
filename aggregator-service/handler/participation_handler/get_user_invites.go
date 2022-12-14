package participationhandler

import (
	"strconv"
	"sync"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/participation"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

func (h *participationHandler) GetUserInvites(c *fiber.Ctx) error {
	uId := c.Params("uid")
	nextPage := c.Query("nextPage")
	limitStr := c.Query("limit")
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	if limit > 40 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 40")
	}

	pi, err := h.participationClient.GetUserInvites(c.Context(), &participation.GetUserInvitesRequest{
		UserId:   uId,
		NextPage: nextPage,
		Limit:    int32(limit),
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	var parties *pbparty.GetManyPartiesMapResponse
	go func() {
		defer wg.Done()
		partyIds := make([]string, len(pi.Invites))
		for i, in := range pi.Invites {
			partyIds[i] = in.PartyId
		}

		parties, _ = h.partyClient.GetManyPartiesMap(c.Context(), &pbparty.GetManyPartiesRequest{Ids: utils.UniqueStringSlice(partyIds)})
	}()

	var inviters *profile.GetManyProfilesMapResponse
	go func() {
		inviterIds := make([]string, len(pi.Invites))
		for i, in := range pi.Invites {
			inviterIds[i] = in.InviterId
		}

		inviters, _ = h.profileClient.GetManyProfilesMap(c.Context(), &profile.GetManyProfilesRequest{Ids: utils.UniqueStringSlice(inviterIds)})
	}()

	wg.Wait()

	aggI := make([]*datastruct.AggregatedPartyInvite, len(pi.Invites))
	for i, pi := range pi.Invites {
		in := datastruct.
			PartyInviteToAgg(pi)

		if inviters.Profiles[pi.UserId] != nil {
			in.AddInviter(datastruct.ProfileToAgg(inviters.Profiles[pi.InviterId]))
		}

		if parties.Parties[pi.PartyId] != nil {
			in.AddParty(datastruct.PartyToAgg(parties.Parties[pi.PartyId]))
		}
		aggI[i] = in
	}

	res := datastruct.PagedAggregatedPartyInvite{
		Invites:  aggI,
		NextPage: pi.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
