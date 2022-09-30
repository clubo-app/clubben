package partyhandler

import (
	"log"
	"sync"

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

	var wg sync.WaitGroup
	wg.Add(3)

	// Get Stories of Party
	go func() {
		defer wg.Done()
		stories, _ := h.storyClient.GetByParty(c.Context(), &sg.GetByPartyRequest{PartyId: party.Id})
		if stories != nil {
			res.AddStory(stories.Stories)
		}
	}()

	// Check if the Requester has already favorited the Party
	go func() {
		defer wg.Done()

		favoriteParty, _ := h.relationClient.GetFavoriteParty(c.Context(), &relation.PartyAndUserRequest{UserId: user.Sub, PartyId: party.Id})
		if favoriteParty != nil {
			res.IsFavorite = true
		}
	}()

	// Check the ParticipationStatus of a User with this Party.
	go func() {
		defer wg.Done()
		if user.Sub != "" {
			participation, err := h.participationClient.GetPartyParticipant(c.Context(), &participation.UserPartyRequest{UserId: user.Sub, PartyId: party.Id})
			if err != nil {
				log.Println("Error getting ParticipationStatus: ", err)
			}
			res.AddParticipationStatus(datastruct.ParseParticipationStatus(participation))
		}
	}()

	wg.Wait()

	return c.Status(fiber.StatusOK).JSON(res)
}
