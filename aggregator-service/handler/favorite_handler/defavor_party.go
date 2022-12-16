package favoritehandler

import (
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h *favoriteHandler) DefavorParty(c *fiber.Ctx) error {
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	pId := c.Params("pId")

	ok, err := h.relationClient.DefavorParty(c.Context(), &rg.PartyAndUserRequest{UserId: user.UserID, PartyId: pId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(ok)
}
