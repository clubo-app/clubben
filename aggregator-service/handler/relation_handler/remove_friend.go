package relationhandler

import (
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h relationGatewayHandler) RemoveFriend(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)
	uId := c.Params("id")

	if user.Sub == uId {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Friend Id")
	}

	ok, err := h.rc.RemoveFriend(c.Context(), &rg.RemoveFriendRequest{UserId: user.Sub, FriendId: uId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(ok)
}
