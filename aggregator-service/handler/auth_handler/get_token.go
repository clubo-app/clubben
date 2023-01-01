package authhandler

import (
	pbauth "github.com/clubo-app/clubben/auth-service/pb/v1"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *authHandler) GetToken(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := h.authClient.CreateToken(c.Context(), &pbauth.CreateTokenRequest{Id: id})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
