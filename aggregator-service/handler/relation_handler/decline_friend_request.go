package relationhandler

import (
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	"github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h relationHandler) DeclineFriendRequest(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)
	fId := c.Params("id")

	if user.Sub == fId {
		return fiber.NewError(fiber.StatusBadRequest, "User id and Friend id are the same")
	}

	ok, err := h.relationClient.DeclineFriendRequest(c.Context(), &relation.DeclineFriendRequestRequest{UserId: user.Sub, FriendId: fId})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	return c.Status(fiber.StatusOK).JSON(ok)
}
