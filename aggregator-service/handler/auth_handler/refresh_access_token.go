package authhandler

import (
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/auth"
	"github.com/gofiber/fiber/v2"
)

func (h authHandler) RefreshAccessToken(c *fiber.Ctx) error {
	rt := c.Params("rt")

	t, err := h.authClient.RefreshAccessToken(c.Context(), &auth.RefreshAccessTokenRequest{RefreshToken: rt})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(t)
}
