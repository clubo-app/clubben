package participationhandler

import (
	"github.com/clubo-app/packages/utils"
	"github.com/clubo-app/packages/utils/middleware"
	"github.com/clubo-app/protobuf/participation"
	"github.com/gofiber/fiber/v2"
)

func (h participationHandler) LeaveParty(c *fiber.Ctx) error {
	pId := c.Params("pid")
	user := middleware.ParseUser(c)

	s, err := h.participationC.LeaveParty(c.Context(), &participation.UserPartyRequest{
		UserId:  user.Sub,
		PartyId: pId,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(s)
}
