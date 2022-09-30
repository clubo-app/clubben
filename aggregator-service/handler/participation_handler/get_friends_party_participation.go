package participationhandler

import (
	"strconv"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h participationHandler) GetFriendsPartyParticipation(c *fiber.Ctx) error {
	uId := c.Params("uid")
	nextPage := c.Query("nextPage")
	limitStr := c.Query("limit")
	limit, _ := strconv.ParseInt(limitStr, 10, 32)
	if limit > 40 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 40")
	}

	friends, err := h.relationClient.GetFriends(c.Context(), &relation.GetFriendsRequest{
		UserId:   uId,
		Limit:    uint64(limit),
		NextPage: nextPage,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	if len(friends.Relations) == 0 {
		return c.Status(fiber.StatusOK).JSON(datastruct.PagedAggregatedParty{Parties: make([]*datastruct.AggregatedParty, 0)})
	}

	friendIds := make([]string, len(friends.Relations))
	for i, friend := range friends.Relations {
		friendIds[i] = friend.FriendId
	}

	participation, err := h.participationClient.GetManyUserParticipations(c.Context(), &participation.GetManyUserParticipationsRequest{
		UserIds: friendIds,
	})

	if len(participation.Participants) == 0 {
		return c.Status(fiber.StatusOK).JSON(datastruct.PagedAggregatedParty{Parties: make([]*datastruct.AggregatedParty, 0)})
	}

	partyIds := make([]string, len(participation.Participants))
	for i, p := range participation.Participants {
		partyIds[i] = p.PartyId
	}

	parties, err := h.partyClient.GetManyParties(c.Context(), &party.GetManyPartiesRequest{
		Ids: partyIds,
	})

	aggParties := make([]*datastruct.AggregatedParty, len(parties.Parties))
	for i, party := range parties.Parties {
		aggParties[i] = datastruct.PartyToAgg(party)
	}

	res := datastruct.PagedAggregatedParty{
		Parties:  aggParties,
		NextPage: friends.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
