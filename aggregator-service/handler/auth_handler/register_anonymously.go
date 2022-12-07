package authhandler

import (
	"github.com/clubo-app/clubben/aggregator-service/datastruct"
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *authHandler) RegisterAnonymously(c *fiber.Ctx) error {
	a, err := h.authClient.RegisterAnonymously(c.Context(), &emptypb.Empty{})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	res := datastruct.LoginResponse{
		Account: datastruct.AccountToAgg(a),
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
