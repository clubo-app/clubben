package participationhandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/gofiber/fiber/v2"
)

func (h participationHandler) JoinParty(c *fiber.Ctx) error {
	pId := c.Params("pid")
	user := middleware.ParseUser(c)

	pp, err := h.participationC.JoinParty(c.Context(), &participation.UserPartyRequest{
		UserId:  user.Sub,
		PartyId: pId,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	res := datastruct.
		PartyParticipantToAgg(pp).
		AddProfile(&datastruct.AggregatedProfile{Id: pp.UserId}).
		AddParty(&datastruct.AggregatedParty{Id: pp.PartyId})

	return c.Status(fiber.StatusCreated).JSON(res)
}
