package partyhandler

import (
	"log"
	"sync"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/clubo-app/clubben/protobuf/relation"
	sg "github.com/clubo-app/clubben/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h *partyHandler) GetParty(c *fiber.Ctx) error {
	id := c.Params("id")
	user, userErr := firebaseauth.GetUser(c)

	party, err := h.partyClient.GetParty(c.Context(), &party.GetPartyRequest{PartyId: id})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profile, _ := h.profileClient.GetProfile(c.Context(), &profile.GetProfileRequest{Id: party.UserId})
	res := datastruct.
		PartyToAgg(party).
		AddCreator(datastruct.ProfileToAgg(profile))

	var wg sync.WaitGroup
	wg.Add(1)

	// Get Stories of Party
	go func() {
		defer wg.Done()
		stories, _ := h.storyClient.GetByParty(c.Context(), &sg.GetByPartyRequest{PartyId: party.Id})
		if stories != nil {
			res.AddStory(stories.Stories)
		}
	}()

	if userErr != nil {
		wg.Add(1)
		// Check if the Requester has already favorited the Party
		go func() {
			defer wg.Done()

			favoriteParty, _ := h.relationClient.GetFavoriteParty(c.Context(), &relation.PartyAndUserRequest{UserId: user.UserID, PartyId: party.Id})
			if favoriteParty != nil {
				res.IsFavorite = true
			}
		}()
	}

	if userErr != nil {
		wg.Add(1)
		// Check the ParticipationStatus of a User with this Party.
		go func() {
			defer wg.Done()
			if user.UserID != "" {
				participation, err := h.participationClient.GetPartyParticipant(c.Context(), &participation.UserPartyRequest{UserId: user.UserID, PartyId: party.Id})
				if err != nil {
					log.Println("Error getting ParticipationStatus: ", err)
				}
				res.AddParticipationStatus(datastruct.ParseParticipationStatus(participation))
			}
		}()
	}

	wg.Wait()

	return c.Status(fiber.StatusOK).JSON(res)
}
