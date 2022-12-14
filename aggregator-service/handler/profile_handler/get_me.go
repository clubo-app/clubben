package profilehandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h profileHandler) GetMe(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	p, err := h.profileClient.GetProfile(c.Context(), &pg.GetProfileRequest{Id: user.Sub})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	if p == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Profile not found")
	}

	res := datastruct.ProfileToAgg(p)

	friendCountRes, _ := h.relationClient.GetFriendCount(c.Context(), &rg.GetFriendCountRequest{UserId: p.Id})
	if friendCountRes != nil {
		res.AddFCount(friendCountRes.FriendCount)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
