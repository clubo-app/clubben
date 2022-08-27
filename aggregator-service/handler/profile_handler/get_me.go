package profilehandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/libs/utils/middleware"
	"github.com/clubo-app/clubben/protobuf/auth"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h profileGatewayHandler) GetMe(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	p, err := h.pc.GetProfile(c.Context(), &pg.GetProfileRequest{Id: user.Sub})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	if p == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Profile not found")
	}

	a, err := h.ac.GetAccount(c.Context(), &auth.GetAccountRequest{Id: user.Sub})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	if a == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Account not found")
	}

	res := datastruct.AggregatedAccount{
		Id:      p.Id,
		Profile: datastruct.ProfileToAgg(p),
		Email:   a.Email,
	}

	friendCountRes, _ := h.rc.GetFriendCount(c.Context(), &rg.GetFriendCountRequest{UserId: p.Id})
	if friendCountRes != nil {
		res.Profile.FriendCount = friendCountRes.FriendCount
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
