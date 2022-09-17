package partyhandler

import (
	"log"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/clubo-app/clubben/protobuf/relation"
	sg "github.com/clubo-app/clubben/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h partyHandler) GetParty(c *fiber.Ctx) error {
	id := c.Params("id")
	user := middleware.ParseUser(c)

	party, err := h.partyClient.GetParty(c.Context(), &party.GetPartyRequest{PartyId: id})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profile, _ := h.profileClient.GetProfile(c.Context(), &profile.GetProfileRequest{Id: party.UserId})
	res := datastruct.
		PartyToAgg(party).
		AddCreator(datastruct.ProfileToAgg(profile))

	stories, _ := h.storyClient.GetByParty(c.Context(), &sg.GetByPartyRequest{PartyId: party.Id})
	if stories != nil {
		res.AddStory(stories.Stories)
	}

	favoriteCount, _ := h.relationClient.GetFavoritePartyCount(c.Context(), &relation.GetFavoritePartyCountRequest{PartyId: party.Id})
	if favoriteCount != nil {
		res.AddFCount(favoriteCount.FavoriteCount)
	}

	if user.Sub != "" {
		participation, err := h.participationClient.GetPartyParticipant(c.Context(), &participation.UserPartyRequest{UserId: user.Sub, PartyId: party.Id})
		if err != nil {
			log.Println("Error getting ParticipationStatus: ", err)
		}
		res.AddParticipationStatus(datastruct.ParseParticipationStatus(participation))
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
