package authhandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	ag "github.com/clubo-app/clubben/protobuf/auth"
	pg "github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

func (h authHandler) Login(c *fiber.Ctx) error {
	req := new(ag.LoginUserRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	a, err := h.authClient.LoginUser(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	p, err := h.profileClient.GetProfile(c.Context(), &pg.GetProfileRequest{
		Id: a.Account.Id,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	res := datastruct.LoginResponse{
		Account: datastruct.AggregatedAccount{
			Id:      a.Account.Id,
			Profile: datastruct.ProfileToAgg(p),
			Email:   a.Account.Email,
		},
		Tokens: *a.Tokens,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
