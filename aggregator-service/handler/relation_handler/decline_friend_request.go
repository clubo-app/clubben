package relationhandler

import (
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h *relationHandler) DeclineFriendRequest(c *fiber.Ctx) error {
	fId := c.Params("id")
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	if user.UserID == fId {
		return fiber.NewError(fiber.StatusBadRequest, "User id and Friend id are the same")
	}

	ok, err := h.relationClient.DeclineFriendRequest(c.Context(), &relation.DeclineFriendRequestRequest{UserId: user.UserID, FriendId: fId})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	return c.Status(fiber.StatusOK).JSON(ok)
}
