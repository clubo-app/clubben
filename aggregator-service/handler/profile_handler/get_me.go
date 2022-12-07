package profilehandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
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

	a, err := h.authClient.GetAccount(c.Context(), &pbauth.GetAccountRequest{Id: user.Sub})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	if a == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Account not found")
	}

	res := datastruct.AccountToAgg(a).AddProfile(datastruct.ProfileToAgg(p))

	friendCountRes, _ := h.relationClient.GetFriendCount(c.Context(), &rg.GetFriendCountRequest{UserId: p.Id})
	if friendCountRes != nil {
		res.Profile.AddFCount(friendCountRes.FriendCount)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
