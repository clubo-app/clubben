package partyhandler

import (
	"log"

	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/clubo-app/clubben/protobuf/relation"
	sg "github.com/clubo-app/clubben/protobuf/story"
	"github.com/gofiber/fiber/v2"
)

func (h partyGatewayHandler) GetParty(c *fiber.Ctx) error {
	log.Println("Running GetParty")
	id := c.Params("id")

	p, err := h.pc.GetParty(c.Context(), &party.GetPartyRequest{PartyId: id})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	profile, _ := h.prf.GetProfile(c.Context(), &profile.GetProfileRequest{Id: p.UserId})

	stories, _ := h.sc.GetByParty(c.Context(), &sg.GetByPartyRequest{PartyId: p.Id})
	favoriteCount, _ := h.rc.GetFavoritePartyCount(c.Context(), &relation.GetFavoritePartyCountRequest{PartyId: p.Id})

	res := datastruct.PartyToAgg(p).AddCreator(datastruct.ProfileToAgg(profile))
	if stories != nil {
		res.AddStory(stories.Stories)
	}
	if favoriteCount != nil {
		res.AddFCount(favoriteCount.FavoriteCount)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
