package participationhandler

import (
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/participation"
	"github.com/gofiber/fiber/v2"
)

func (h participationHandler) DeclinePartyInvite(c *fiber.Ctx) error {
	pId := c.Params("pid")
	uId := c.Params("uid")
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	s, err := h.participationClient.DeclinePartyInvite(c.Context(), &participation.DeclinePartyInviteRequest{
		UserId:    user.UserID,
		PartyId:   pId,
		InviterId: uId,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(s)
}
