package relationhandler

import (
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h relationHandler) CreateFriendRequest(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	fId := c.Params("id")

	if user.Sub == fId {
		return fiber.NewError(fiber.StatusBadRequest, "You can't add yourself")
	}

	ok, err := h.relationClient.CreateFriendRequest(c.Context(), &relation.CreateFriendRequestRequest{UserId: user.Sub, FriendId: fId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(ok)
}
