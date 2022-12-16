package participationhandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/gofiber/fiber/v2"
)

func (h *participationHandler) AcceptPartyInvite(c *fiber.Ctx) error {
	pId := c.Params("pid")
	uId := c.Params("uid")
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	pp, err := h.participationClient.AcceptPartyInvite(c.Context(), &participation.DeclinePartyInviteRequest{
		UserId:    user.UserID,
		PartyId:   pId,
		InviterId: uId,
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
