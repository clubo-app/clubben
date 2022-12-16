package profilehandler

import (
	"github.com/clubo-app/clubben/libs/utils"
	"github.com/clubo-app/clubben/protobuf/profile"
	"github.com/gofiber/fiber/v2"
)

type UsernameTakenResponse struct {
	Taken bool `json:"taken"`
}

func (h *profileHandler) UsernameTaken(c *fiber.Ctx) error {
	uName := c.Params("username")

	res, err := h.profileClient.UsernameTaken(c.Context(), &profile.UsernameTakenRequest{Username: uName})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(&UsernameTakenResponse{Taken: res.Taken})
}
