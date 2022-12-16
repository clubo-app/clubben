package relationhandler

import (
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h *relationHandler) AcceptFriendRequest(c *fiber.Ctx) error {
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	fId := c.Params("id")

	if user.UserID == fId {
		return fiber.NewError(fiber.StatusBadRequest, "You can't accept Requests from yourself")
	}

	ok, err := h.relationClient.AcceptFriendRequest(c.Context(), &relation.AcceptFriendRequestRequest{UserId: user.UserID, FriendId: fId})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	return c.Status(fiber.StatusOK).JSON(ok)
}
