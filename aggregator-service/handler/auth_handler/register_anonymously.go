package authhandler

import (
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RegisterAnonymouslyResponse struct {
	CustomToken string `json:"custom_token"`
}

func (h *authHandler) RegisterAnonymously(c *fiber.Ctx) error {
	token, err := h.authClient.RegisterAnonymously(c.Context(), &emptypb.Empty{})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	res := RegisterAnonymouslyResponse{
		CustomToken: token.CustomToken,
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
