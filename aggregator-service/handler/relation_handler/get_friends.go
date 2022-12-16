package relationhandler

import (
	"strconv"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h relationHandler) GetFriends(c *fiber.Ctx) error {
	user, userErr := firebaseauth.GetUser(c)

	uId := c.Params("id")
	nextPage := c.Query("nextPage")

	acceptedStr := c.Query("accepted")
	accepted, acceptedErr := strconv.ParseBool(acceptedStr)

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	if limit > 40 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 40")
	}

	// find the friends or incoming friend requests of the wanted user
	fr := new(rg.PagedFriendRelations)
	if !accepted && acceptedErr == nil {
		var err error
		fr, err = h.relationClient.GetIncomingFriendRequests(c.Context(), &rg.GetIncomingFriendRequestsRequest{UserId: uId, NextPage: nextPage, Limit: limit})
		if err != nil {
			return utils.ToHTTPError(err)
		}
	} else {
		var err error
		fr, err = h.relationClient.GetFriends(c.Context(), &rg.GetFriendsRequest{UserId: uId, NextPage: nextPage, Limit: limit})
		if err != nil {
			return utils.ToHTTPError(err)
		}
	}

	ids := make([]string, len(fr.Relations))
	for i, fp := range fr.Relations {
		ids[i] = fp.FriendId
	}

	// we get the profile of all the user having a relation to the wanted user
	profiles, err := h.profileClient.GetManyProfilesMap(c.Context(), &profile.GetManyProfilesRequest{Ids: utils.UniqueStringSlice(ids)})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profileRes := make([]*datastruct.AggregatedProfile, len(profiles.Profiles))
	for i, f := range fr.Relations {
		p := profiles.Profiles[f.FriendId]
		aggP := datastruct.ProfileToAgg(p)

		if userErr != nil {
			// we see if the friends are also freinds of the requester, this way we can display if the user is already friends with the friends
			aggP.AddFs(datastruct.ParseFriendShipStatus(user.UserID, f))
		}

		profileRes[i] = aggP
	}

	res := datastruct.PagedAggregatedProfile{
		Profiles: profileRes,
		NextPage: fr.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
