package authhandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	firebaseauth "github.com/clubo-app/clubben/libs/firebase-auth"
	"github.com/clubo-app/clubben/libs/utils"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	rg "github.com/clubo-app/clubben/protobuf/relation"
	"github.com/gofiber/fiber/v2"
)

func (h *authHandler) GetMe(c *fiber.Ctx) error {
	user, userErr := firebaseauth.GetUser(c)
	if userErr != nil {
		return userErr
	}

	a, err := h.authClient.GetAccount(c.Context(), &pbauth.GetAccountRequest{
		Id: user.UserID,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	p, err := h.profileClient.GetProfile(c.Context(), &pg.GetProfileRequest{Id: user.UserID})
	if err != nil {
		return utils.ToHTTPError(err)
	}
	if p == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Profile not found")
	}

	res := datastruct.AccountToAgg(a).AddProfile(datastruct.ProfileToAgg(p))

	friendCountRes, _ := h.relationClient.GetFriendCount(c.Context(), &rg.GetFriendCountRequest{UserId: p.Id})
	if friendCountRes != nil {
		res.Profile.AddFCount(friendCountRes.FriendCount)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
