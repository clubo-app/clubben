package authhandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	ag "github.com/clubo-app/clubben/protobuf/auth"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

func (h authHandler) GoogleLogin(c *fiber.Ctx) error {
	token := c.Params("token")

	account, err := h.authClient.GoogleLoginUser(c.Context(), &ag.GoogleLoginUserRequest{
		Token: token,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	p, err := h.profileClient.GetProfile(c.Context(), &pg.GetProfileRequest{
		Id: account.Account.Id,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	res := datastruct.LoginResponse{
		Account: datastruct.AggregatedAccount{
			Id:      account.Account.Id,
			Profile: datastruct.ProfileToAgg(p),
			Email:   account.Account.Email,
		},
		Tokens: *account.Tokens,
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
