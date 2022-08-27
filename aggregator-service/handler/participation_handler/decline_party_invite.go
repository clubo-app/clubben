package participationhandler

import (
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	"github.com/clubo-app/protobuf/participation"
	"github.com/gofiber/fiber/v2"
)

func (h participationHandler) DeclinePartyInvite(c *fiber.Ctx) error {
	pId := c.Params("pid")
	uId := c.Params("uid")
	user := middleware.ParseUser(c)

	s, err := h.participationC.DeclinePartyInvite(c.Context(), &participation.DeclinePartyInviteRequest{
		UserId:    user.Sub,
		PartyId:   pId,
		InviterId: uId,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(s)
}
