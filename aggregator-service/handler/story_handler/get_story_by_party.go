package storyhandler

import (
	"strconv"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	"github.com/clubo-app/clubben/protobuf/profile"
	sg "github.com/clubo-app/clubben/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h *storyHandler) GetStoryByParty(c *fiber.Ctx) error {
	pId := c.Params("id")

	limitStr := c.Query("limit")
	limit, _ := strconv.ParseUint(limitStr, 10, 32)
	if limit > 20 {
		return fiber.NewError(fiber.StatusBadRequest, "Max limit is 20")
	}

	nextPage := c.Query("nextPage")

	stories, err := h.storyClient.GetByParty(c.Context(), &sg.GetByPartyRequest{
		PartyId:  pId,
		NextPage: nextPage,
		Limit:    uint32(limit),
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	party, _ := h.partyClient.GetParty(c.Context(), &pbparty.GetPartyRequest{
		PartyId: pId,
	})

	// Get all the ids of all story creators
	ids := make([]string, len(stories.Stories))
	for i, s := range stories.Stories {
		ids[i] = s.UserId
	}

	// here we aggregate all the profiles of the story creators
	profilesRes, _ := h.profileClient.GetManyProfilesMap(c.Context(), &profile.GetManyProfilesRequest{Ids: utils.UniqueStringSlice(ids)})

	aggS := make([]*datastruct.AggregatedStory, len(stories.Stories))
	for i, story := range stories.Stories {
		s := datastruct.
			StoryToAgg(story).
			AddParty(datastruct.PartyToAgg(party))

		if profilesRes != nil {
			s = s.AddCreator(datastruct.ProfileToAgg(profilesRes.Profiles[story.UserId]))
		}
		aggS[i] = s
	}

	res := datastruct.PagedAggregatedStory{
		Stories:  aggS,
		NextPage: stories.NextPage,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
