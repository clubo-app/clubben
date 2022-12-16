package relationhandler

import (
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h *relationHandler) RemoveFriend(c *fiber.Ctx) error {
	uId := c.Params("id")
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	if user.UserID == uId {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Friend Id")
	}

	ok, err := h.relationClient.RemoveFriend(c.Context(), &rg.RemoveFriendRequest{UserId: user.UserID, FriendId: uId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(ok)
}
