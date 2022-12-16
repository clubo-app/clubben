package relationhandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h relationHandler) CreateFriendRequest(c *fiber.Ctx) error {
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	fId := c.Params("id")

	if user.UserID == fId {
		return fiber.NewError(fiber.StatusBadRequest, "You can't add yourself")
	}

	fr, err := h.relationClient.CreateFriendRequest(c.Context(), &relation.CreateFriendRequestRequest{UserId: user.UserID, FriendId: fId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(datastruct.ParseFriendShipStatus(user.UserID, fr))
}
