package partyhandler

import (
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/party"
	"github.com/gofiber/fiber/v2"
)

func (h *partyHandler) DeleteParty(c *fiber.Ctx) error {
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	pId := c.Params("id")

	res, err := h.partyClient.DeleteParty(c.Context(), &party.DeletePartyRequest{RequesterId: user.UserID, PartyId: pId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
