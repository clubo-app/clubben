package participationhandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/gofiber/fiber/v2"
)

func (h participationHandler) AcceptPartyInvite(c *fiber.Ctx) error {
	pId := c.Params("pid")
	uId := c.Params("uid")
	user := middleware.ParseUser(c)
	pp, err := h.participationC.AcceptPartyInvite(c.Context(), &participation.DeclinePartyInviteRequest{
		UserId:    user.Sub,
		PartyId:   pId,
		InviterId: uId,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	res := datastruct.
		PartyParticipantToAgg(pp).
		AddUser(datastruct.AggregatedProfile{Id: pp.UserId}).
		AddParty(datastruct.AggregatedParty{Id: pp.PartyId})

	return c.Status(fiber.StatusCreated).JSON(res)
}
