package authhandler

import (
	"github.com/clubo-app/clubben/libs/utils"
	ag "github.com/clubo-app/clubben/protobuf/auth"
	"github.com/gofiber/fiber/v2"
)

func (h authHandler) GoogleLogin(c *fiber.Ctx) error {
	token := c.Params("token")

	res, err := h.authClient.GoogleLoginUser(c.Context(), &ag.GoogleLoginUserRequest{
		Token: token,
	})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
