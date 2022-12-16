package participationhandler

import (
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/gofiber/fiber/v2"
)

func (h *participationHandler) LeaveParty(c *fiber.Ctx) error {
	pId := c.Params("pid")
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	s, err := h.participationClient.LeaveParty(c.Context(), &participation.UserPartyRequest{
		UserId:  user.UserID,
		PartyId: pId,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(s)
}
