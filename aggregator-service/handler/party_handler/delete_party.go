package partyhandler

import (
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	pbparty "github.com/clubo-app/clubben/party-service/pb/v1"
	"github.com/gofiber/fiber/v2"
)

func (h *partyHandler) DeleteParty(c *fiber.Ctx) error {
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	pId := c.Params("id")

	res, err := h.partyClient.DeleteParty(c.Context(), &pbparty.DeletePartyRequest{RequesterId: user.UserID, PartyId: pId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
